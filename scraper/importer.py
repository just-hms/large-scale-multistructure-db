from pymongo import MongoClient
from pymongo.errors import DuplicateKeyError

from faker import Faker

import json

#Type hinting imports
from datetime import datetime
from typing import Literal


#######
#UTILS#
#######

##############
#DB FUNCTIONS#
##############

def makeUser(usersCollection,userName:str,type:Literal["user","barber","admin"])->int:
    """
    Function used to generate a new user. Returns the newly inserted document id.
    If a user with the same username already existed, return -1.
    """

    #Create the user dict structure
    user = {}
    user["username"] = userName
    user["email"] = f"{userName}@barbershop.com"
    user["password"] = f"{userName}1234"
    user["type"] = type
    user["currentAppointment"] = {}

    #Add user to the db and return its new id
    try:
        return usersCollection.insert_one(user).inserted_id
    except DuplicateKeyError:
        return -1

def addAppointmentToUser(usersCollection,userId,shopId:str,shopName:str,startDate:datetime,duration:datetime):
    """Adds current appointment info to the specified user"""

    #Create the appointment dict structure
    appointment = {}
    appointment["shopId"] = shopId
    appointment["shopName"] = shopName
    appointment["startDate"] = startDate
    appointment["duration"] = duration

    #Update the specified user's appointment
    usersCollection.update_one({
        "_id": userId
    },{
        "$set": {"currentAppointment":appointment}
    })

def makeShop(shopsCollection,shopData:dict)->int:
    """Function used to generate a new barber shop. Uses the data format from the scraper. Returns the newly inserted document id."""

    #Create a new temp object to do what we need without dirtying the one we got passed
    shop = shopData.copy()

    #Modify the shop dict to suit our needs
    ## Fix phone field
    if shop["phone"] == []:
        shop["phone"] = ""
    ##Add reviews once the shop exists. Delete them in the meanwhile.
    shop.pop("reviewData",None)
    ##Prepare coordinates key better
    lat = shop["coordinates"].split(" ")[0]
    lon = shop["coordinates"].split(" ")[1]
    shop["coordinates"] = {}
    shop["coordinates"]["lat"] = float(lat)
    shop["coordinates"]["lon"] = float(lon)
    ##Prepare fields
    shop["appointments"] = []
    shop["views"] = []
    shop["reviews"] = []

    #Add shop to the db and return its new id
    return shopsCollection.insert_one(shop).inserted_id

def addReviewToShop(shopsCollection,shopId,userId,shopReview):
    """Adds a review to the specified shop. Uses the data format from the scraper."""

    #Create the review dict structure
    review = {}
    review["userId"] = userId
    review["username"] = shopReview["username"]
    review["rating"] = shopReview["rating"]
    review["reported"] = False
    review["content"] = shopReview["body"]
    #We generate a review date as we do not have it
    review["createdAt"] = Faker().date_time_between(start_date='-10y', end_date='now').strftime("%d/%m/%Y, %H:%M")

    #Update the specified barber shop's review list
    shopsCollection.update_one({
        "_id": shopId
    },{
        "$push": {"reviews":review}
    })


def main():

    #Establish connection to databases
    mongoClient = MongoClient('localhost', 27017)

    #TODO: Remove after testing
    #Reset databases
    mongoClient.drop_database("barberShop")

    #Connect to Mongo instance and create database and collections if needed
    barberDatabaseMongo = mongoClient["barberShop"]
    usersCollectionMongo = barberDatabaseMongo["users"]
    barberShopsCollectionMongo = barberDatabaseMongo["barberShops"]

    #Make usernames unique
    usersCollectionMongo.create_index("username",unique=True)

    #Load scraped data
    scrapedData = {}
    with open("aioScrapingResults.json","r") as scrapedDataFile:
        scrapedData = json.load(scrapedDataFile)

    #TODO:Test, remove later
    shop = scrapedData["Roma"][0]
    shopId = makeShop(barberShopsCollectionMongo,shop)
    for review in shop["reviewData"]["reviews"]:
        userId = makeUser(usersCollectionMongo,review["username"],"user")
        if userId != -1:
            addReviewToShop(barberShopsCollectionMongo,shopId,userId,review)
    cursor = barberShopsCollectionMongo.find({})
    for doc in cursor:
        print(doc)

    #Go through the scraped data, location by location
    #for location, shopsList in scrapedData.items():
    #    for shop in shopsList:
    #        #Make the document structure the way we want to




if __name__ == "__main__":
    main()