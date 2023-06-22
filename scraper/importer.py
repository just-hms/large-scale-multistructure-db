from pymongo import MongoClient, GEOSPHERE
from pymongo.errors import DuplicateKeyError

from faker import Faker

from datetime import timedelta
from datetime import datetime
import json
import random
import time

import bcrypt
import uuid

import argparse

#Type hinting imports
from typing import Literal

fake = Faker()

#####################
# UTILITY FUNCTIONS #
#####################

def roundUpDateTime(date:datetime, delta:timedelta):
    """Round up a datetime to the nearest timedelta provided."""
    return date + (datetime.min - date) % delta


##############
#DB FUNCTIONS#
##############

def makeUser(usersCollection,userName:str,type:Literal["user","barber","admin"])->int:
    """
    Function used to generate a new user. Returns the newly inserted document id.
    If a user with the same username already existed, return -1.
    """

    #Clean the username first as it might contain spaces
    userName = userName.replace(" ", "")
    #Create the user dict structure
    user = {}
    user["_id"] = str(uuid.uuid4())
    user["username"] = userName
    user["email"] = f"{userName}@barbershop.com"
    user["password"] = f"{userName}1234"
    #Hash and salt password
    user["password"] = bcrypt.hashpw(user["password"].encode('utf-8'), bcrypt.gensalt(12)).decode("utf-8")
    user["type"] = type
    user["signupDate"] = fake.date_time_between(start_date='-10y', end_date='now')
    user["ownedShops"] = []
    user["currentAppointment"] = {}

    #Add user to the db and return its new id
    try:
        return usersCollection.insert_one(user).inserted_id, user
    except DuplicateKeyError:
        return -1, None

def addAppointmentToUser(usersCollection,userId,appointment):
    """Adds current appointment info to the specified user"""

    #Update the specified user's appointment
    usersCollection.update_one({
        "_id": userId
    },{
        "$set": {"currentAppointment":appointment}
    })

def fakeBarber(usersCollection):
    """Makes a barber user with a faked name. Returns the newly inserted document id."""
    fakeUserName = fake.simple_profile()["username"]
    return makeUser(usersCollection,fakeUserName,"barber")

def addOwnedShopToBarber(usersCollection,userId,shopId):
    """Adds the specified shop to the list of shops owned by a barber."""

    #Check that the user is a barber first
    user = usersCollection.find_one({"_id":userId})
    if user["type"] != "barber":
        print("Only barbers can own shops!")
        return

    usersCollection.update_one({
        "_id": userId
    },{
        "$push": {"ownedShops":shopId}
    })

def makeShop(shopsCollection,shopData:dict)->int:
    """Function used to generate a new barber shop. Uses the data format from the scraper. Returns the newly inserted document id."""

    #Create a new temp object to do what we need without dirtying the one we got passed
    shop = shopData.copy()

    #Modify the shop dict to suit our needs
    shop["_id"] = str(uuid.uuid4())
    shop["description"] = f"Welcome to {shop['name']}"
    ## Fix phone field
    if shop["phone"] == []:
        shop["phone"] = ""
    ## Fix null ratings
    shop["rating"] = shop["rating"] if shop["rating"] > 0 else 1
    ##Add reviews once the shop exists. Delete them in the meanwhile.
    shop.pop("reviewData",None)
    ##Remove calendar
    shop.pop("calendar")
    ##Rename location field
    shop["address"] = shop["location"]
    shop.pop("location")
    ##Prepare coordinates for Mongo usage
    lat = float(shop["coordinates"].split(" ")[0])
    lon = float(shop["coordinates"].split(" ")[1])
    shop.pop("coordinates")
    shop["location"] = {}
    shop["location"]["type"] = "Point"
    shop["location"]["coordinates"] = [lon,lat]
    ##Fake number of employees
    shop["employees"] = random.randint(1,3)

    #Add shop to the db and return its new id
    return shopsCollection.insert_one(shop).inserted_id

def addReviewToShop(reviewsCollection,shopId,user,shopReview,upvotesIdList,downvotesIdList):
    """Adds a review to the specified shop. Uses the data format from the scraper."""

    #Create the review dict structure
    review = {}
    #Generate an id for the review
    review["_id"] = str(uuid.uuid4())
    review["shopId"] = shopId
    review["userId"] = user["_id"]
    review["username"] = shopReview["username"].replace(" ", "")
    review["rating"] = shopReview["rating"] if shopReview["rating"] > 0 else 1
    review["reported"] = False
    review["content"] = shopReview["body"]
    review["upvotes"] = upvotesIdList
    review["downvotes"] = downvotesIdList
    #We generate a review date as we do not have it
    review["createdAt"] = fake.date_time_between(start_date=user["signupDate"], end_date='now')

    #Update the review collection
    reviewsCollection.insert_one(review)

def addViewsToShop(shopviewsCollectionMongo,viewsList):
    """Adds a list of view info to the specified shop"""

    #Insert the views list in its collection
    shopviewsCollectionMongo.insert_many(viewsList)

def addAppointmentsToShop(appointmentsCollectionMongo,shopAppointmentsList):
    """Adds list of appointments to the specified shop"""

    #Insert the appointments list in its collection
    appointmentsCollectionMongo.insert_many(shopAppointmentsList)

def fakeViews(shopviewsCollectionMongo,shopId,userList,maxViewsAmount=1500):
    """Generate fake views up to maxAmount. Needs an array of userIds to choose from. 
        Returns array of generated userId-creationDate pairs."""

    #Get a random amount 
    viewsAmount = random.randint(1,maxViewsAmount)

    viewsUserList = []
    for _ in range(viewsAmount):
        #Extract a user
        user = random.choice(userList)
        #Fake view date
        view = {}
        view["_id"] = str(uuid.uuid4())
        view["createdAt"] = fake.date_time_between(start_date=user["signupDate"], end_date='now')
        view["userId"] = user["_id"]
        view["shopId"] = shopId
        viewsUserList.append(view)

    #Add data to the DB
    addViewsToShop(shopviewsCollectionMongo,viewsUserList)

    #Return view date
    return viewsUserList

def fakeAppointments(usersCollection,appointmentsCollectionMongo,shopId,shopName,viewsList,generatedUsersMap,maxAppointmentsAmount=200):
    """Generate fake appointments up to maxAmount. Needs an array of {userId,viewDate} to choose from."""

    #Get a random amount 
    appointmentsAmount = random.randint(1,maxAppointmentsAmount)
    #Set a normalized cancellation probability
    cancelProbability = 0.05

    #Prepare list of data to be inserted in the DB
    shopAppointmentsList = []
    for _ in range(appointmentsAmount):
        randomView = random.choice(viewsList)
        appointment = {}
        #Add id to appointment
        appointment["_id"] = str(uuid.uuid4())
        #Fake appointment date
        appointment["createdAt"] = fake.date_time_between(start_date=randomView["createdAt"], end_date=randomView["createdAt"]+timedelta(minutes=5))
        appointment["startDate"] = fake.date_time_between(start_date=appointment["createdAt"], end_date=appointment["createdAt"]+timedelta(days=5))
        #Round datetime to the nearest half hour
        appointment["startDate"] = roundUpDateTime(appointment["startDate"],timedelta(minutes=30))
        #Fake appointment status, with a small chance to cancel it
        if appointment["startDate"] > datetime.utcnow():
            appointment["status"] = "pending"
        else:
            appointment["status"] = "completed"
        if random.random() < cancelProbability:
            appointment["status"] = "canceled"
        #Add shopID
        appointment["shopId"] = shopId
        appointment["shopName"] = shopName
        addAppointmentToUser(usersCollection,randomView["userId"],appointment)
        #Fill appointment info for the shop
        appointment["userId"] = randomView["userId"]
        appointment["username"] = generatedUsersMap[randomView["userId"]]["username"]
        shopAppointmentsList.append(appointment)

    #Add data to the DB
    addAppointmentsToShop(appointmentsCollectionMongo,shopAppointmentsList)

def fakeUserList(userList,maxAmount=50):
    """Pull a random list of at max maxAmount users. Used to create upvotes and downvotes lists."""

    #Get minimum bewteen amount of users and maxAmount
    maxAmount = len(userList) if len(userList)<maxAmount else maxAmount
    #Get a random amount 
    amount = random.randint(1,maxAmount)

    #Extract the user ids
    return random.choices(userList,k=amount) 



def main():

    mongoHost = '127.0.0.1'
    mongoPort = 27017

    #Parse the command-line arguments
    argParser = argparse.ArgumentParser()
    argParser.add_argument("-H", "--host", type=str, help="The IP address of the machine hosting the MongoDB instance")
    argParser.add_argument("-P", "--port", type=int, help="The port of the machine hosting the MongoDB instance")

    args = argParser.parse_args()

    if args.host:
        mongoHost = args.host

    if args.port:
        mongoPort = args.port

    print(f"{mongoHost}:{mongoPort}")

    start_time = time.perf_counter()
    print("> Starting BarberShop importer\n")

    #Establish connection to databases
    mongoClient = MongoClient(mongoHost,mongoPort)

    #Reset databases
    mongoClient.drop_database("barbershop")
    mongoClient.drop_database("barberShop")

    #Connect to Mongo instance and create database and collections if needed
    barberDatabaseMongo = mongoClient["barbershop"]
    usersCollectionMongo = barberDatabaseMongo["users"]
    barberShopsCollectionMongo = barberDatabaseMongo["barbershops"]
    shopviewsCollectionMongo = barberDatabaseMongo["shopviews"]
    appointmentsCollectionMongo = barberDatabaseMongo["appointments"]
    reviewsCollectionMongo = barberDatabaseMongo["reviews"]

    #Make username and email unique
    usersCollectionMongo.create_index("username",unique=True)
    usersCollectionMongo.create_index("email",unique=True)
    #Prepare Mongo for geolocation
    barberShopsCollectionMongo.create_index([("location",GEOSPHERE)])

    #Load scraped data
    scrapedData = {}
    with open("aioScrapingResults.json","r") as scrapedDataFile:
        scrapedData = json.load(scrapedDataFile)

    #Prepare useful data structures
    generatedShopsIds = []
    generatedUsers = {}
    importedShops = 0

    #Add the Admin to database
    makeUser(usersCollectionMongo,"admin","admin")

    #Go through the scraped data, location by location
    for _, shopsList in scrapedData.items():
        for shop in shopsList:
            print(f">> Importing {shop['name']}")
            importedShops += 1
            #Generate a shop entry in the database
            shopId = makeShop(barberShopsCollectionMongo,shop)
            generatedShopsIds.append(shopId)
            #Generate a fake barber user for the shop and save its id.
            #We might accidentally generate a barber with the same username. Repeat until we succeed.
            barberId = -1
            while True:
                barberId, _ = fakeBarber(usersCollectionMongo)
                if barberId != -1:
                    break
            #Add shop to list of shops owned by the barber
            addOwnedShopToBarber(usersCollectionMongo,barberId,shopId)
            #Go through the reviews and generate users based on the found usernames. Skip if username already exists.
            for review in shop["reviewData"]["reviews"]:
                userId, user = makeUser(usersCollectionMongo,review["username"],"user")
                while userId == -1:
                    review["username"] = review["username"] + "1"
                    userId, user = makeUser(usersCollectionMongo,review["username"],"user")
                generatedUsers[userId] = user
                #Add review to shop while faking amount of upvotes and downvotes
                addReviewToShop(reviewsCollectionMongo,shopId,user,review,fakeUserList(list(generatedUsers.keys())),fakeUserList(list(generatedUsers.keys()),5))
            ##Fake interaction stuff we do not have: Views, Appointments

            #Fake a random amount of views from random users. Max 1500.
            viewsUserList = fakeViews(shopviewsCollectionMongo,shopId,list(generatedUsers.values()),1500)
            #Fake a random number of appointments
            fakeAppointments(usersCollectionMongo,appointmentsCollectionMongo,shopId,shop["name"],viewsUserList,generatedUsers,200)

    #Print results
    end_time = time.perf_counter()
    print(f"\n## Imported {importedShops} shops in {end_time - start_time} seconds")
    print(f"## Database statistics:\n")
    print(barberDatabaseMongo.command("dbstats"))





if __name__ == "__main__":
    main()