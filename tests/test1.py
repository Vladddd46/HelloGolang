import requests 
import json
import time

# Config
block_number = 11509788
domain       = "localhost:8080"
##############

i = 0
while(i < block_number + 100):
	url      = "http://{domain}/api/block/{block_number}/total".format(domain = domain, block_number = block_number + i)
	response = requests.get(url).text
	serealized_data = json.loads(response)
	print("transactions: " + str(serealized_data["transactions"]) + " total: " + str(serealized_data["amount"]))
	i += 1
	time.sleep(1) # API can be used 5 times/second, thus delay in 1 sec per request is used.