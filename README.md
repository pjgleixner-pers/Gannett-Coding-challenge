# Gannett-Coding-challenge

## Acceptance Criteria
 * (✔️) The produce database is a single, in memory array of data and supports reads and writes
 * (✔️) Supports adding more than one new produce item at a time
 * (✔️) The produce includes name, produce code, and unit price
 * (✔️) The produce name is alphanumeric and case insensitive
 * (✔️) The produce codes are sixteen characters long, with dashes separating each four character group
 * (✔️) The produce codes are alphanumeric and case insensitive
 * (✔️) The produce unit price is a number with up to 2 decimal places
 * (✔️) Error handling (GET nonexistent produce)
 * (✔️) Error handling (bad POST payload) 
 * (✔️) The API is RESTful
 * (✔️) The functionality is tested and verifiable without manually exercising the endpoints
 * (✔️) The API supports adding and deleting individual produce. You can also get any produce in the database.
 * (❗) Produce Code is unique (no apple or pear have the same code)

## Running the program
Make sure you have golang installed (https://go.dev/doc/install)
### By default
1. In terminal, go to the project folder
2. run the command : "go run main.go"
3. to exit use Ctrl + C
### Using VSCode
Make sure you have installed the Go extension
1. Go to main.go
2. select Run -> Start debuging (or F5 if using the IntelliJ IDEA Keybindings extension)
    or
2. select Run -> Run without debuging (or Shift + F9 if using the IntelliJ IDEA Keybindings extension)
3. to stop the program on top should appear a red square (🟥)

## Starting Data
```
[
    {
        "ID": "A12T-4GH7-QPL9-3N4M",
        "Name": "Lettuce",
        "Price": "3.46"
    },
    {
        "ID": "E5T6-9UI3-TH15-QR88",
        "Name": "Peach",
        "Price": "2.99"
    },
    {
        "ID": "YRT6-72AS-K736-L4AR",
        "Name": "Green Pepper",
        "Price": "0.79"
    },
    {
        "ID": "TQ4C-VV6T-75ZX-1RMR",
        "Name": "Gala Apple",
        "Price": "3.59"
    }
]
```
## Supported Endpoints
### GET
#### Show all Items:
* Endpoint: 
    /item
* Example Request:
    curl --request GET --url localhost:8080/item
* Expected Response:
    [{"ID":"A12T-4GH7-QPL9-3N4M","Name":"Lettuce","Price":"3.46"},{"ID":"E5T6-9UI3-TH15-QR88","Name":"Peach","Price":"2.99"},{"ID":"YRT6-72AS-K736-L4AR","Name":"Green Pepper","Price":"0.79"},{"ID":"TQ4C-VV6T-75ZX-1RMR","Name":"Gala Apple","Price":"3.59"}]

#### Show item by ID:
* Endpoint: 
    /item/{id}
* Example Request:
    curl --request GET --url localhost:8080/item/A12T-4GH7-QPL9-3N4M
* Expected Response:
    {"ID":"A12T-4GH7-QPL9-3N4M","Name":"Lettuce","Price":"3.46"}
#### Error handling:
If it is run with an Id that does not exist in the data will trow a 404 Not Found
Example:
item id = 1234-1234-1234-1234 doesn't exist in this example
Request: 
    curl -i --request GET --url localhost:8080/item/1234-1234-1234-1234
Respponse:
    HTTP/1.1 404 Not Found
    Date: Thu, 27 Jan 2022 01:21:29 GMT
    Content-Length: 31
    Content-Type: text/plain; charset=utf-8
    
    {"ID":"","Name":"","Price":""}
### POST
#### Single Post:
* Endpoint: 
    /items
* Example Request:
    curl -i -X POST -H "Content-type:application/json" --data "{\"ID\":\"QWER-1234-TYUI-5678\",\"Name\":\"Apple\",\"Price\":\"1.99\"}" localhost:8080/items
* Expected Response:
    HTTP/1.1 201 Created
    Date: Thu, 27 Jan 2022 02:12:00 GMT
    Content-Length: 61
    Content-Type: text/plain; charset=utf-8

    [{"ID":"QWER-1234-TYUI-5678","Name":"Apple","Price":"1.99"}]
#### Multiple Post:
* Endpoint: 
    /items
* Example Request:
    curl -i -X POST -H "Content-type:application/json" --data "[{\"ID\":\"QWER-1234-TYUI-5678\",\"Name\":\"Apple\",\"Price\":\"1.99\"},{\"ID\":\"ZXCV-5678-VBNM-9012\",\"Name\":\"Tomato\",\"Price\":\"0.99\"}]" localhost:8080/items
* Expected Response:
   HTTP/1.1 201 Created
    Date: Thu, 27 Jan 2022 02:19:00 GMT
    Content-Length: 121
    Content-Type: text/plain; charset=utf-8

    [{"ID":"QWER-1234-TYUI-5678","Name":"Apple","Price":"1.99"},{"ID":"ZXCV-5678-VBNM-9012","Name":"Tomato","Price":"0.99"}]     
#### Error handling:
If you send wrong data (ID with no format "AAAA-AAAA-AAAA-AAAA" or Price with no two decimal places) will trow a 400 Bad Request
Example:
Request: 
    curl -i -X POST -H "Content-type:application/json" --data "{\"ID\":\"QWER-1234-TYUI\",\"Name\":\"Apple\",\"Price\":\"1.99\"}" localhost:8080/items
Respponse: 
    HTTP/1.1 400 Bad Request
    Date: Thu, 27 Jan 2022 02:24:40 GMT
    Content-Length: 0

### DELETE
* Endpoint: 
    /items/{id}
* Example Request:
    curl -i --request DELETE --url localhost:8080/items/A12T-4GH7-QPL9-3N4M
* Expected Response:
    HTTP/1.1 204 No Content
#### Error handling:
If it is run with an Id that does not exist in the data will trow a 404 Not Found
Example:
item id = 1234-1234-1234-1234 doesn't exist in this example
Request: 
    curl -i --request DELETE --url localhost:8080/items/1234-1234-1234-1234
Respponse:
    HTTP/1.1 404 Not Found
    Date: Thu, 27 Jan 2022 01:30:09 GMT
    Content-Length: 0

## Testing
If you are using VSCode go to the file main_test.go. There will be the test functions TestShowItems, TestShowItemById, TestDeleteItem, TestCreateItem and TestCreateItems. Above each one will be the option to *run test* or *debug test*. All test in the initial state should pass unless you change the *req, err* variable or the *itemPayload* variable.


