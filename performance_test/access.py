import requests
import json

for i in range(10000):
    data = {
        'clientid': 'b48f5b98-e9e7-40ce-b8cf-cdc4d2c59061',
        'password': 'pass'
    }
    headers = {'Content-type': 'application/json', 'Accept': 'text/plain'}
    r = requests.post('http://localhost:8080/addlocker',
                      data=json.dumps(data), headers=headers)
    print(r)
