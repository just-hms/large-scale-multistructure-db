import googlemaps

import json
import requests
import time

API_KEY = "[API_KEY_HERE]"

scrapingResults = {}

def getPlacePhoto(photoReference):

    url = f"https://maps.googleapis.com/maps/api/place/photo?maxwidth=1600&photo_reference={photoReference}&key={API_KEY}"
    response = requests.request("GET", url)

    return response.url

def main(locationsList = ["Roma","Firenze","Milano", "Palermo", "New York"]):

    with open("mapsApiKey.txt","r") as file:
        API_KEY = file.readline()

    gmaps = googlemaps.Client(key=API_KEY)

    with open("mapsScrapingResults.json","r+") as file:
        #Load previously scraped data
        scrapingResults = json.load(file)

        #Go through each specified location
        for location in locationsList:

            #Init data structures on first visit to location
            if not location in scrapingResults:
                scrapingResults[location] = {}
                scrapingResults[location]["scrapedShopsNames"] = []
                scrapingResults[location]["scrapedShopsData"] = []
            #Init next page token
            nextPageToken = ""

            #Fetch info from Yelp for a given city
            barberListRaw = gmaps.places(query=f"barbers in {location}", radius=20000)
            #Loop all the pages for a city until results are available starting from here
            while True:
                #Remember nextPageToken if present 
                if "next_page_token" in barberListRaw:
                    nextPageToken = barberListRaw["next_page_token"]
                else:
                    nextPageToken = ""

                for barberShop in barberListRaw["results"]:
                    #Skip barber if its data has already been fetched and saved
                    if barberShop["name"] in scrapingResults[location]["scrapedShopsNames"]:
                        print(f'{barberShop["name"]} in {location} was already scraped. Skipping...')
                        continue
                    #Format info on a barberShop
                    print(f'Fetching data on {barberShop["name"]} in {location}')
                    shopData = {}
                    shopData["name"] = barberShop["name"]
                    shopData["rating"] = barberShop["rating"]
                    shopData["location"] = barberShop["formatted_address"]
                    shopData["coordinates"] = f'{barberShop["geometry"]["location"]["lat"]} {barberShop["geometry"]["location"]["lng"]}'
                    #TODO: Places Photo request necessary?
                    #shopData["imageLink"] = barberShop["photos"][0]["html_attributions"][0]
                    shopData["imageLink"] = ""

                    #Get shop details 
                    shopDetails = gmaps.place(place_id=barberShop["place_id"])["result"]
                    if "international_phone_number" in shopDetails:
                        shopData["phone"] = shopDetails["international_phone_number"]
                    else:
                        shopData["phone"] = []

                    #Get shop calendar
                    shopData["calendar"] = []
                    if "opening_hours" in shopDetails:
                        for slot in shopDetails["opening_hours"]["periods"]:
                            calendarSlot = {}
                            calendarSlot["is_overnight"] = False
                            calendarSlot["start"] = slot["open"]["time"]
                            calendarSlot["end"] = slot["close"]["time"]
                            calendarSlot["day"] = slot["open"]["day"]
                            shopData["calendar"].append(calendarSlot)


                    #Get shop reviews
                    shopData["reviewData"] = {}
                    shopData["reviewData"]["reviews"] = []
                    for review in shopDetails["reviews"]:
                        reviewData = {}
                        reviewData["username"] = review["author_name"]
                        reviewData["rating"] = review["rating"]
                        reviewData["body"] = review["text"]
                        shopData["reviewData"]["reviews"].append(reviewData)

                    #Save data to file
                    scrapingResults[location]["scrapedShopsNames"].append(shopData['name'])
                    scrapingResults[location]["scrapedShopsData"].append(shopData)
                    file.seek(0)
                    json.dump(scrapingResults,file)

                #Check if there are any more results to fetch. Next tokens become valid after a bit they are issued, so keep trying.
                if nextPageToken == "":
                    break
                else:
                    while True:
                        try:
                            barberListRaw = gmaps.places(page_token=nextPageToken)
                            break
                        except:
                            time.sleep(1)
                




if __name__ == "__main__":
    main()