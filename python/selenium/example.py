from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By

chrome_options = Options()
chrome_options.add_argument('--headless')
chrome_options.add_argument('--no-sandbox')
chrome_options.add_argument('--disable-dev-shm-usage')

browser = webdriver.Chrome(chrome_options=chrome_options, executable_path="./chromedriver")

browser.get("https://giovannirossini.github.io")

element = browser.find_element(By.CLASS_NAME, 'btn-contact').click()

print(element)