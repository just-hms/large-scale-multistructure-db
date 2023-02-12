import googlemaps

import json
import requests
import time

API_KEY = "[API_KEY_HERE]"

scrapingResults = {}

def getPlacePhotoUrl(key,photoReference):

    url = f"https://maps.googleapis.com/maps/api/place/photo?maxwidth=1600&photo_reference={photoReference}&key={key}"
    response = requests.request("GET", url)

    return response.url

def main(locationsList = ["Roma","Firenze","Milano", "Palermo", "New York"],needsPhotoUrlPatch=False):

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
                        #Check if we need to add previously skipped image data to the entry. This can be slow.
                        if needsPhotoUrlPatch and "photos" in barberShop:
                            for scrapedShop in scrapingResults[location]["scrapedShopsData"]:
                                if scrapedShop["name"] != barberShop["name"]:
                                    continue
                                if scrapedShop["imageLink"] == "":
                                    scrapedShop["imageLink"] = getPlacePhotoUrl(key=API_KEY,photoReference=barberShop["photos"][0]["photo_reference"])
                                    print(f'Adding photo to {barberShop["name"]} in {location}...')
                                else:
                                    print(f'{barberShop["name"]} in {location} from Maps was already scraped. Skipping...')
                                break
                        else:
                            print(f'{barberShop["name"]} in {location} from Maps was already scraped. Skipping...')
                        continue

                    #Format info on a barberShop
                    print(f'Fetching data on {barberShop["name"]} in {location} from Maps')
                    shopData = {}
                    shopData["name"] = barberShop["name"]
                    shopData["rating"] = barberShop["rating"]
                    shopData["location"] = barberShop["formatted_address"]
                    shopData["coordinates"] = f'{barberShop["geometry"]["location"]["lat"]} {barberShop["geometry"]["location"]["lng"]}'
                    if "photos" in barberShop:
                        shopData["imageLink"] = getPlacePhotoUrl(key=API_KEY,photoReference=barberShop["photos"][0]["photo_reference"])
                    else:
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
                            if "close" in slot:
                                calendarSlot["end"] = slot["close"]["time"]
                            else:
                                calendarSlot["end"] = ""
                            calendarSlot["day"] = slot["open"]["day"]
                            shopData["calendar"].append(calendarSlot)


                    #Get shop reviews
                    shopData["reviewData"] = {}
                    shopData["reviewData"]["reviews"] = []
                    if "review" in shopDetails:
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