from selenium import webdriver
from webdriver_manager.chrome import ChromeDriverManager
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.common.action_chains import ActionChains
from selenium.webdriver.support import expected_conditions as EC
from selenium.common.exceptions import ElementNotVisibleException
from selenium.webdriver.common.by import By
from selenium.common.exceptions import NoSuchElementException
  
import time

# As there are possibilities of different chrome
# browser and we are not sure under which it get
# executed let us use the below syntax
driver = webdriver.Chrome(ChromeDriverManager().install())

driver.maximize_window()
#driver.implicitly_wait(30)

url = "https://www.google.com/maps/search/barber+in+pisa"
driver.get(url)

wait = WebDriverWait(driver, 5)

wait.until(EC.element_to_be_clickable(
    (By.XPATH, "//button[contains(@aria-label, 'Rifiuta tutto')]"))).click()

barberShopsData = []
barberNameList = []
barberListUnique = []
barberListRaw = driver.find_elements(By.XPATH, "//div[@role='article']")

#Make the barber list unique
for barber in barberListRaw: 
    shopName = barber.get_attribute("aria-label")
    if shopName not in barberNameList: 
        barberListUnique.append(barber) 
        barberNameList.append(shopName)

for shop in barberListUnique:
    shopData = {}
    shopData['name'] = shop.get_attribute("aria-label")
    shop.click()
    wait.until(EC.text_to_be_present_in_element((By.CSS_SELECTOR, ".fontHeadlineLarge"),shopData['name']))
    try:
        shopData["rating"] = driver.find_element_by_xpath('//*[@id="QA0Szd"]/div/div/div[1]/div[3]/div/div[1]/div/div/div[2]/div[2]/div[1]/div[1]/div[2]/div/div[1]/div[2]/span[1]/span/span[1]').text
    except NoSuchElementException:
        shopData["rating"] = -1
    shopData["location"] = driver.find_element_by_xpath("//button[contains(@aria-label, 'Indirizzo:')]").get_attribute("aria-label").split("Indirizzo: ")[1]
    try:
        shopData["phone"] = driver.find_element_by_xpath("//button[contains(@aria-label, 'Telefono:')]").get_attribute("aria-label").split("Telefono: ")[1]
    except NoSuchElementException:
        shopData["phone"] = -1
    shopData["calendar"] = []
    #Expand calendar
    try:
        driver.find_element_by_xpath("//div[contains(@jsaction, 'pane.openhours')]").click()
        calendar = wait.until(EC.visibility_of_element_located((By.XPATH, "//div[contains(@aria-label, 'Nascondi orari')]"))).find_elements_by_tag_name("tr")
        for calendarSlot in calendar:
            shopData["calendar"].append(calendarSlot.text)
    except NoSuchElementException:
        pass
    barberShopsData.append(shopData)
print(barberShopsData)

driver.quit()
