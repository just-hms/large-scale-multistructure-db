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
    user["ownedShops"] = []
    user["currentAppointment"] = {}

    #Add user to the db and return its new id
    try:
        return usersCollection.insert_one(user).inserted_id
    except DuplicateKeyError:
        return -1

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

def addOwnedShopToBarber(shopsCollection,usersCollection,userId,shopId):
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
    ##Prepare fields
    shop["appointments"] = []
    shop["views"] = []
    shop["reviews"] = []

    #Add shop to the db and return its new id
    return shopsCollection.insert_one(shop).inserted_id

def addReviewToShop(shopsCollection,shopId,userId,shopReview,upvotesIdList,downvotesIdList):
    """Adds a review to the specified shop. Uses the data format from the scraper."""

    #Create the review dict structure
    review = {}
    #Generate an id for the review
    review["_id"] = str(uuid.uuid4())
    review["userId"] = userId
    review["username"] = shopReview["username"].replace(" ", "")
    review["rating"] = shopReview["rating"]
    review["content"] = shopReview["body"]
    review["upvotes"] = upvotesIdList
    review["downvotes"] = downvotesIdList
    #We generate a review date as we do not have it
    review["createdAt"] = fake.date_time_between(start_date='-10y', end_date='now')

    #Update the specified barber shop's review list
    shopsCollection.update_one({
        "_id": shopId
    },{
        "$push": {"reviews":review}
    })

def addViewsToShop(shopsCollection,shopId,viewsList):
    """Adds a list of view info to the specified shop"""

    #Update the specified shop's views
    shopsCollection.update_one({
        "_id": shopId
    },{
        "$set": {"views":viewsList}
    })

def addAppointmentsToShop(shopsCollection,shopId,shopAppointmentsList):
    """Adds current appointment info to the specified user"""

    #Update the specified user's appointment
    shopsCollection.update_one({
        "_id": shopId
    },{
        "$set": {"appointments":shopAppointmentsList}
    })

def fakeViews(shopsCollection,shopId,userList,maxViewsAmount=1500):
    """Generate fake views up to maxAmount. Needs an array of userIds to choose from. 
        Returns array of generated userId-creationDate pairs."""

    #Get a random amount 
    viewsAmount = random.randint(1,maxViewsAmount)

    viewsUserList = []
    for _ in range(viewsAmount):
        randomUserId = random.choice(userList)
        #Fake view date
        viewDate = fake.date_time_between(start_date='-10y', end_date='now')
        viewsUserList.append({"userId":randomUserId,"viewCreation":viewDate})

    #Add data to the DB
    addViewsToShop(shopsCollection,shopId,viewsUserList)

    #Return view date
    return viewsUserList

def fakeAppointments(usersCollection,shopsCollection,shopId,shopName,viewsList,maxAppointmentsAmount=200):
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
        appointment["createdAt"] = fake.date_time_between(start_date=randomView["viewCreation"], end_date=randomView["viewCreation"]+timedelta(minutes=5))
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
        #Make a copy to be used for users
        userAppointment = appointment.copy()
        userAppointment["shopId"] = shopId
        userAppointment["shopName"] = shopName
        addAppointmentToUser(usersCollection,randomView["userId"],userAppointment)
        #Fill appointment info for the shop
        appointment["userId"] = randomView["userId"]
        shopAppointmentsList.append(appointment)

    #Add data to the DB
    addAppointmentsToShop(shopsCollection,shopId,shopAppointmentsList)

def fakeUserList(userList,maxAmount=50):
    """Pull a random list of at max maxAmount users. Used to create upvotes and downvotes lists."""

    #Get minimum bewteen amount of users and maxAmount
    maxAmount = len(userList) if len(userList)<maxAmount else maxAmount
    #Get a random amount 
    amount = random.randint(1,maxAmount)

    #Extract the user ids
    return random.choices(userList,k=amount) 



def main():
    start_time = time.perf_counter()
    print("> Starting BarberShop importer\n")

    #Establish connection to databases
    mongoClient = MongoClient('localhost', 27017)

    #Reset databases
    mongoClient.drop_database("barbershop")

    #Connect to Mongo instance and create database and collections if needed
    barberDatabaseMongo = mongoClient["barbershop"]
    usersCollectionMongo = barberDatabaseMongo["users"]
    barberShopsCollectionMongo = barberDatabaseMongo["barbershops"]

    #Make usernames unique
    usersCollectionMongo.create_index("username",unique=True)
    #Prepare Mongo for geolocation
    barberShopsCollectionMongo.create_index([("location",GEOSPHERE)])

    #Load scraped data
    scrapedData = {}
    with open("aioScrapingResults.json","r") as scrapedDataFile:
        scrapedData = json.load(scrapedDataFile)

    #Prepare useful data structures
    generatedShopsIds = []
    generatedUsersIds = []
    importedShops = 0

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
                barberId = fakeBarber(usersCollectionMongo)
                if barberId != -1:
                    break
            #Add shop to list of shops owned by the barber
            addOwnedShopToBarber(barberShopsCollectionMongo,usersCollectionMongo,barberId,shopId)
            #Go through the reviews and generate users based on the found usernames. Skip if username already exists.
            for review in shop["reviewData"]["reviews"]:
                userId = makeUser(usersCollectionMongo,review["username"],"user")
                if userId != -1:
                    generatedUsersIds.append(userId)
                    #Add review to shop while faking amount of upvotes and downvotes
                    addReviewToShop(barberShopsCollectionMongo,shopId,userId,review,fakeUserList(generatedUsersIds),fakeUserList(generatedUsersIds,5))
            ##Fake interaction stuff we do not have: Views, Appointments

            #Fake a random amount of views from random users. Max 1500.
            viewsUserList = fakeViews(barberShopsCollectionMongo,shopId,generatedUsersIds,1500)
            #Fake a random number of appointments
            fakeAppointments(usersCollectionMongo,barberShopsCollectionMongo,shopId,shop["name"],viewsUserList,200)

    #Print results
    end_time = time.perf_counter()
    print(f"\n## Imported {importedShops} shops in {end_time - start_time} seconds")
    print(f"## Database statistics:\n")
    print(barberDatabaseMongo.command("dbstats"))





if __name__ == "__main__":
    main()