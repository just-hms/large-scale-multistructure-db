from yelpapi import YelpAPI

import json

API_KEY = "IIQPcbU9wwdg83l3iRYVqXO-LSEw1yoAXD8680iv80cD4TQzbT0vKv9HwWh5p8qA_b3o8XqF_k6_Re0ytgFEV0VT8Wmobs8W0CjzVFEUvgfx2dlY9clMIpQBErzSY3Yx"

scrapingResults = {}

def main(locationsList = ["Roma","Firenze","Milano", "Palermo", "New York"]):

    with open("yelpScrapingResults.json","r+") as file:
        with YelpAPI(API_KEY) as yelp_api:

            #Load previously scraped data
            scrapingResults = json.load(file)

            #Go through each specified location
            for location in locationsList:

                #Init data structures on first visit to location
                if not location in scrapingResults:
                    scrapingResults[location] = {}
                    scrapingResults[location]["scrapedShopsNames"] = []
                    scrapingResults[location]["scrapedShopsData"] = []

                #Fetch info from Yelp for a given city
                barberListRaw = yelp_api.search_query(categories='barbers', location=location, locale="it_IT",  limit=50, radius=20000)

                for barberShop in barberListRaw["businesses"]:
                    #Skip barber if its data has already been fetched and saved
                    if barberShop["name"] in scrapingResults[location]["scrapedShopsNames"]:
                        print(f'{barberShop["name"]} in {location} was already scraped. Skipping...')
                        continue
                    #Format info on a barberShop
                    print(f'Fetching data on {barberShop["name"]} in {location}')
                    shopData = {}
                    shopData["name"] = barberShop["name"]
                    shopData["rating"] = barberShop["rating"]
                    shopData["location"] = f'{barberShop["location"]["address1"]}, {barberShop["location"]["zip_code"]} {barberShop["location"]["city"]} {barberShop["location"]["state"]}'
                    shopData["coordinates"] = f'{barberShop["coordinates"]["latitude"]} {barberShop["coordinates"]["longitude"]}'
                    shopData["phone"] = barberShop["phone"]
                    shopData["imageLink"] = barberShop["image_url"]

                    #Get shop calendar
                    shopDetails = yelp_api.business_query(id=barberShop["id"])
                    if "hours" in shopDetails:
                        shopData["calendar"] = shopDetails["hours"][0]["open"]
                    else:
                        shopData["calendar"] = []

                    #Get shop reviews
                    shopReviews = yelp_api.reviews_query(id=barberShop["id"])
                    shopData["reviewData"] = {}
                    shopData["reviewData"]["reviews"] = []
                    for review in shopReviews["reviews"]:
                        reviewData = {}
                        reviewData["username"] = review["user"]["name"]
                        reviewData["rating"] = review["rating"]
                        reviewData["body"] = review["text"]
                        shopData["reviewData"]["reviews"].append(reviewData)

                    #Save data to file
                    scrapingResults[location]["scrapedShopsNames"].append(shopData['name'])
                    scrapingResults[location]["scrapedShopsData"].append(shopData)
                    file.seek(0)
                    json.dump(scrapingResults,file)




if __name__ == "__main__":
    main()