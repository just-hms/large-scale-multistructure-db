import json

import yelpApiScraper
import mapsApiScraper

def main():
    print("> Starting All-In-One Scraper")
    #Set desired scraping targets
    locations = ["Roma","Firenze","Milano", "Palermo", "New York"]
    print("> Scraping the following target locations: "+str(locations))
    #Launch the various independent scrapers
    print(">> Starting YELP scraper")
    yelpApiScraper.main(locations)
    print(">> YELP scraper done")
    print(">> Starting GMaps scraper")
    mapsApiScraper.main(locations)
    print(">> GMaps scraper done")

    #Fetch the results from the various files
    scrapingResults = {}
    yelpResults = {}
    mapsResults = {}
    with open("yelpScrapingResults.json","r") as yelpResultsFile:
        with open("mapsScrapingResults.json","r") as mapsResultsFile:
            yelpResults = json.load(yelpResultsFile)
            mapsResults = json.load(mapsResultsFile)

            print("> Putting results together")
            for location in locations:
                scrapingResults[location] = []
                cityBarberListYelp = yelpResults[location]["scrapedShopsData"]
                cityBarberListMaps = mapsResults[location]["scrapedShopsData"]

                print(f">> Fixing {location}")
                #Fetch everything we can from Maps first
                for barberShopMaps in cityBarberListMaps:
                    print(f">>> Found from Maps: {barberShopMaps['name']} in {location}")
                    #If the place also exists in the Yelp list, add their reviews together
                    if barberShopMaps["name"] in yelpResults[location]["scrapedShopsNames"]:
                        for barberShopYelp in cityBarberListYelp:
                            if barberShopYelp["name"] == barberShopMaps["name"]:
                                for review in barberShopYelp["reviewData"]["reviews"]:
                                    barberShopMaps["reviewData"]["reviews"].append(review)
                                print(f">>>> Found Yelp reviews of {barberShopMaps['name']}")
                                #While we are at it, add a photo if it is not present
                                if barberShopMaps["imageLink"] == "":
                                    barberShopMaps["imageLink"] = barberShopYelp["imageLink"]
                                break
                            else:
                                continue
                    
                    scrapingResults[location].append(barberShopMaps)

                #Now check if Yelp has some places that Maps does not have
                for barberShopYelp in cityBarberListYelp:
                    if not barberShopYelp["name"] in mapsResults[location]["scrapedShopsNames"]:
                        print(f">>> Found from Yelp: {barberShopYelp['name']} in {location}")
                        scrapingResults[location].append(barberShopYelp)

    #Save the compound results into a file
    print("> Saving results")
    with open("aioScrapingResults.json","w") as aioResultsFile:
        json.dump(scrapingResults,aioResultsFile)





if __name__ == "__main__":
    main()