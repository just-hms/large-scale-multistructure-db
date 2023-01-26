from selenium import webdriver
from webdriver_manager.chrome import ChromeDriverManager
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.common.action_chains import ActionChains
from selenium.webdriver.support import expected_conditions as EC
from selenium.common.exceptions import ElementNotVisibleException
from selenium.webdriver.common.by import By
from selenium.common.exceptions import NoSuchElementException
from selenium.common.exceptions import TimeoutException

import json


# As there are possibilities of different chrome
# browser and we are not sure under which it get
# executed let us use the below syntax
driver = webdriver.Chrome(ChromeDriverManager().install())
driver.maximize_window()

wait = WebDriverWait(driver, 10)
loadingWait = WebDriverWait(driver, 120)
actions = ActionChains(driver)

scrapingResults = {}

def main():

    firstTime = True
    locationsList = ["pisa"]
    alreadyDiscoveredBarbers = {}

    with open("mapsScrapingResults.json","r+") as file:
        #TODO:Load previously fetched data to resume scraping if stopped
        #scrapingResults = json.load(file)
        for location in locationsList:
            url = "https://www.google.com/maps/search/barbieri+a+"+location
            driver.get(url)

            if firstTime:
                wait.until(EC.element_to_be_clickable(
                    (By.XPATH, "//button[contains(@aria-label, 'Rifiuta tutto')]"))).click()
                firstTime = False

            scrapingResults[location] = {}
            scrapingResults[location]["scrapedShopsNames"] = []
            scrapingResults[location]["scrapedShopsData"] = []
            barberNameList = []
            barberListUnique = []
            barberListRaw = driver.find_elements(By.XPATH, "//div[@role='article']")

            #Make the barber list unique
            for barber in barberListRaw: 
                shopName = barber.get_attribute("aria-label")
                if shopName not in barberNameList: 
                    barberListUnique.append(barber) 
                    barberNameList.append(shopName)

            while len(barberListUnique) > 0:
                shop = barberListUnique.pop(0)
                shopData = {}
                #Get barberShop name to detect when its info menu successfully expands
                shopData['name'] = shop.get_attribute("aria-label")
                #TODO:Skip if already previously parsed
                #if shopName in 
                print("Fetching info from: "+shopData['name'])
                #Expand barberShop info. Try until it succeeds as a loading might be happening
                while True:
                    try:
                        print("Clicking on: "+shopData['name'])
                        shop.click()
                        print("Waiting expansion on: "+shopData['name'])
                        wait.until(EC.text_to_be_present_in_element((By.CSS_SELECTOR, ".fontHeadlineLarge"),shopData['name']))
                        print("Expanded: "+shopData['name'])
                        #Get barberShop into view
                        driver.execute_script("arguments[0].scrollIntoView();", shop)
                        #If the loading spinner became visible, wait until it disappears
                        loadingWait.until(EC.invisibility_of_element_located((By.XPATH, "//div[contains(@jspan, 't-WPtQSFf6msE')]")))
                        #Fetch barberShop info
                        try:
                            shopData["rating"] = driver.find_element_by_xpath('//*[@id="QA0Szd"]/div/div/div[1]/div[3]/div/div[1]/div/div/div[2]/div[2]/div[1]/div[1]/div[2]/div/div[1]/div[2]/span[1]/span/span[1]').text
                        except NoSuchElementException:
                            shopData["rating"] = -1
                        try:
                            shopData["location"] = driver.find_element_by_xpath("//button[contains(@aria-label, 'Indirizzo:')]").get_attribute("aria-label").split("Indirizzo: ")[1]
                        except NoSuchElementException:
                            shopData["location"] = -1
                        try:
                            shopData["phone"] = driver.find_element_by_xpath("//button[contains(@aria-label, 'Telefono:')]").get_attribute("aria-label").split("Telefono: ")[1]
                        except NoSuchElementException:
                            shopData["phone"] = -1
                        try:
                            shopData["imageLink"] = driver.find_element_by_xpath("//button[contains(@aria-label, 'Foto di ')]/img").get_attribute("src")
                        except NoSuchElementException:
                            shopData["imageLink"] = -1
                        break
                    except TimeoutException:
                        continue
                shopData["calendar"] = []
                #Expand barberShop calendar if present. Try until it succeeds as a loading might be happening
                while True:
                    try:
                        driver.find_element_by_xpath("//div[contains(@jsaction, 'pane.openhours')]").click()
                        calendar = wait.until(EC.visibility_of_element_located((By.XPATH, "//div[contains(@aria-label, 'Nascondi orari')]"))).find_elements_by_tag_name("tr")
                        for calendarSlot in calendar:
                            shopData["calendar"].append(calendarSlot.text)
                        break
                    except NoSuchElementException:
                        break
                    except TimeoutException:
                        continue
                
                ###############
                ### REVIEWS ###
                ###############
                shopData["reviewData"] = {}
                shopData["reviewData"]["reviews"] = []
                reviewsList = []
                reviewsIDList = []
                foundReviewsList = []
                #Expand review info if present. Try until it succeeds as a loading might be happening
                print("Expanding reviews")
                while True:
                    try:
                        driver.find_element_by_xpath("//*[@id='QA0Szd']/div/div/div[1]/div[3]/div/div[1]/div/div/div[2]/div[2]/div[1]/div[1]/div[2]/div/div[1]/div[2]/span[2]").click()
                        break
                    except NoSuchElementException:
                        break
                    except TimeoutException:
                        continue
                #Wait until a review is visible
                wait.until(EC.element_to_be_clickable((By.XPATH, "//button[contains(@aria-label, 'Azioni per la recensione')]")))
                #Check if new reviews became visible
                foundReviewsList = driver.find_elements(By.XPATH, "//button[contains(@aria-label, 'Azioni per la recensione')]/../../../../..")
                for review in foundReviewsList: 
                    reviewID = review.get_attribute("review-id")
                    if reviewID not in reviewsIDList: 
                        reviewsList.append(review) 
                        reviewsIDList.append(reviewID)
                #Scroll review info if present. Try until it succeeds as a loading might be happening
                while len(reviewsList) > 0:
                    review = reviewsList.pop(0)
                    reviewData = {}
                    while True:
                        try:
                            #Fetch review info
                            reviewData["username"] = review.get_attribute("aria-label")
                            print("Got username: "+str(reviewData["username"]))
                            try:
                                reviewData["rating"] = review.find_element_by_xpath(".//span[contains(@class, 'kvMYJc')]").get_attribute("aria-label").split(" ")[1]
                            except NoSuchElementException:
                                reviewData["rating"] = -1
                            print("Got rating: "+str(reviewData["rating"]))
                            #Expand review if needed, then wait until it does
                            print("Expanding review: "+review.get_attribute("aria-label"))
                            try:
                                wait.until(EC.element_to_be_clickable((By.XPATH, ".//button[contains(@aria-label, 'Mostra di pi')]"))).click()
                                #driver.find_element_by_xpath(".//button[contains(@aria-label, 'Mostra di pi')]").click()
                                #If there is a need to expand, the review element must be refreshed
                                #review = driver.find_element_by_xpath(f".//div[contains(@aria-label, '{reviewData.username}')]")
                            except TimeoutException:
                                pass
                            try:
                                reviewData["body"] = review.find_element_by_xpath(".//div[contains(@class, 'wiI7pd')]").text
                            except NoSuchElementException:
                                reviewData["body"] = -1
                            print("Got body: "+str(reviewData["body"]))
                            #Scroll to review
                            driver.execute_script("arguments[0].scrollIntoView();", review)
                            break
                        except NoSuchElementException:
                            break
                        except TimeoutException:
                            continue

                    #Append review data
                    shopData["reviewData"]["reviews"].append(reviewData)
                    #Check if new reviews became visible
                    foundReviewsList = driver.find_elements(By.XPATH, "//div[contains(@jsan, 't-4VhXSdTzr88')]")
                    for review in foundReviewsList: 
                        reviewID = review.get_attribute("review-id")
                        if reviewID not in reviewsIDList: 
                            reviewsList.append(review) 
                            reviewsIDList.append(reviewID)

                ##############
                ### SAVING ###
                ##############
                #Append data
                scrapingResults[location]["scrapedShopsNames"].append(shopData['name'])
                scrapingResults[location]["scrapedShopsData"].append(shopData)
                #Save data to file and overwrite evry time
                file.seek(0)
                json.dump(scrapingResults,file)

                #Check if new barberShops became visible
                barberListRaw = driver.find_elements(By.XPATH, "//div[@role='article']")
                for barber in barberListRaw: 
                    shopName = barber.get_attribute("aria-label")
                    if shopName not in barberNameList: 
                        barberListUnique.append(barber) 
                        barberNameList.append(shopName)

        print(scrapingResults)

    driver.quit()

if __name__ == "__main__":
    main()


    #t-WPtQSFf6msE,7.lXJj5c,7.Hk4XGb,t-6j59QVuhtOE
    #t-WPtQSFf6msE,7.lXJj5c,7.Hk4XGb,t-6j59QVuhtOE
