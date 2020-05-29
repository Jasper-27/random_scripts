from urllib.request import urlopen as uReq
import re

#Getting the reddit profile
username = "root_27"
reddit = "https://www.reddit.com/user/"
URL = reddit + username

#Tries to connect 10 times before giving up
tries = 0
maxTries = 10

while tries < maxTries:
    try:
        #Gets the page from the URL
        uClient = uReq(URL)
        page_html = uClient.read()
        uClient.close()

        #Looks through the page for matches to a regex expressions
        pattern = '\"totalKarma\":(.*?),'
        result = re.findall(pattern, str(page_html))[0]
        print (result)

        break
    except:
        tries += 1

        if tries > maxTries - 1:
            print("Can't connect to Reddit right now")
            quit()
