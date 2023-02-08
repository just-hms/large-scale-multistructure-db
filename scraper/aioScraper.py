import json

import yelpApiScraper
import mapsApiScraper

def main():
    yelpError = False
    mapsError = False
    print("> Starting All-In-One Scraper")
    #Set desired scraping targets
    locations = ["Roma","Firenze","Milano","Palermo","Aquila","Potenza","Catanzaro","Napoli","Bologna","Trieste","Genova","Ancona",
                "Torino","Bari","Cagliari","Trento","Perugia","Aosta","Venezia","Los Angeles","Boston","Detroit","Washington",
                "Tampa","Houston","Phoenix","Denver","Honolulu","Columbus","Richmond","New York","Paris", "Berlin", "Madrid",
                "London","Vienna","Amsterdam", "Brussels", "Lisbon", "Prague","Copenhagen","Stockholm", "Oslo", "Helsinki", 
                "Reykjavik", "Dublin", "Bratislava", "Ljubljana", "Zagreb", "Sarajevo", "Belgrade", "Skopje", "Athens","Valletta", "Chisinau", 
                "Monaco", "Andorra la Vella", "San Marino", "Vatican City",
                "Beijing", "Bangkok", "Jakarta", "New Delhi", "Tokyo", "Seoul", "Manila",    
                "Mumbai", "Shanghai", "Karachi", "Istanbul", "Tehran", "Baghdad",    
                "Riyadh", "Singapore", "Kuala Lumpur", "Abu Dhabi", "Doha", "Jerusalem",    
                "Baku", "Muscat", "Kuwait City", "Astana", "Tashkent", "Damascus",    
                "Sana'a", "Dhaka", "Cairo", "Amman", "Ankara", "Tbilisi", "Colombo",    
                "Bishkek", "Phnom Penh", "Vientiane", "Kathmandu", "Ulaanbaatar", "Hanoi",    
                "Vientiane", "Kathmandu", "Pyongyang", "Islamabad", "Dili", "Thimphu",
                "Montgomery", "Juneau", "Little Rock", "Sacramento",    
                "Hartford", "Dover", "Tallahassee", "Atlanta",    
                "Boise", "Springfield", "Indianapolis", "Des Moines",    
                "Topeka", "Frankfort", "Baton Rouge", "Augusta", "Annapolis",    
                "Lansing", "St. Paul", "Jackson", "Jefferson City",    
                "Helena", "Lincoln", "Carson City", "Concord", "Trenton",    
                "Santa Fe", "Albany", "Raleigh", "Bismarck",
                "Oklahoma City", "Harrisburg", "Providence", "Columbia",    
                "Pierre", "Nashville", "Austin", "Salt Lake City", "Montpelier",    
                "Richmond", "Olympia", "Cheyenne", "Madison",
                "Birmingham", "Bradford", "Bristol", "Cambridge", "Canterbury",    
                "Carlisle", "Chester", "Chichester", "Coventry", "Derby",    
                "Durham", "Ely", "Exeter", "Gloucester", "Hereford", "Kingston upon Hull",    
                "Lancaster", "Leeds", "Leicester", "Lichfield", "Lincoln", "Liverpool",    
                "Manchester", "Newcastle upon Tyne", "Norwich", "Nottingham",    
                "Oxford", "Peterborough", "Plymouth", "Portsmouth", "Preston", "Ripon",    
                "Salford", "Salisbury", "Sheffield", "Southampton", "St Albans",    
                "Stoke-on-Trent", "Sunderland", "Truro", "Wakefield", "Wells",    
                "Westminster", "Winchester", "Wolverhampton", "Worcester", "York"]

    print("> Scraping the following target locations: "+str(locations))
    #Launch the various independent scrapers
    print(">> Starting YELP scraper")
    try:
        yelpApiScraper.main(locations)
    except:
        print(">> WARNING: Yelp has errored. It is highly likely that the data rate was exceeded.")
        yelpError = True
    print(">> YELP scraper done")
    print(">> Starting GMaps scraper")
    try:
        mapsApiScraper.main(locations)
    except:
        print(">> WARNING: Maps has errored. It is highly likely that the data rate was exceeded.")
        mapsError = True
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

    if yelpError:
        print(">> WARNING: Yelp has errored. It is highly likely that the data rate was exceeded.")
    if mapsError:
        print(">> WARNING: Maps has errored. It is highly likely that the data rate was exceeded.")




if __name__ == "__main__":
    main()