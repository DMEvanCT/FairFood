#Fair Food Service 

This service is used to get and update fair food for any given fair. 
It allows people to search for food by type and update food as they add
it to the fair


In order to understand the request / return values please see the proto files

* WhatKindofFood.proto 


###To build for ALPINE and build docker container
```bash 
GOOS=linux GOARCH=amd64 go build -o fairfood
docker build . -t (tag)
```

###To build for local machine to test
```bash 
go build ./... -o fairfood
```




To test on your own 
 * Create a digitial ocean account using the link 
 https://m.do.co/c/72fbe213f818
 * Spin up a mysql instance. A base instance will work just fine. 
 * Apply the demo sql located in this repo to the sql instance
 * Spin up a small kubernetes cluster 
    * kubectl create namespace fair
 * Create the FairFood.yaml file 
```yaml 
FoodService:
  username: (username
  password: (password)
  host:(sql address)
  port: (sql port)
```
  * Run kubectl create secret generic foodservice --from-file=(FairService.yaml)
  * Run kubectl apply -f fairfood.yaml 

 


In order to test you can use evans CLI located at 
https://github.com/ktr0731/evans


You are more than welcome to test this our while it is up 


Desert Example
```json 
WhatKindofFood.FoodInMyBelly@134.209.142.86:50051> call GetFood

foodtype (TYPE_STRING) => Desert
{
  "foodname": "Fried Dough",
  "foodprice": 5.5,
  "fooddesc": "Dough that has been fried with powdered sugar or cinamon",
  "vendor": "Food For Fun",
  "foodloc": "Whites Farm Road",
  "foodtype": "Desert"
}
{
  "foodname": "Fried Oreos",
  "foodprice": 6,
  "fooddesc": "Oreos that have been fried",
  "vendor": "Food For Fun",
  "foodloc": "Whites Farm Road",
  "foodtype": "Desert"
}
```
Dinner Example
```json 
foodtype (TYPE_STRING) => Dinner
{
  "foodname": "Cheeseburger Sliders",
  "foodprice": 8.7,
  "fooddesc": "Oreos that have been fried",
  "vendor": "Harrys Sliders",
  "foodloc": "Canfield  Road",
  "foodtype": "Dinner"
}
{
  "foodname": "Corn Dogs",
  "foodprice": 3.5,
  "fooddesc": "Fried hot dogs on a stick",
  "vendor": "Durham Elks Club",
  "foodloc": "Canfield  Road",
  "foodtype": "Dinner"
}
{
  "foodname": "Pulled Pork",
  "foodprice": 11,
  "fooddesc": "Smoked pulled pork sandwitch topped with pickles ",
  "vendor": "Porkys Barbecue",
  "foodloc": "Town Hall  Road",
  "foodtype": "Dinner"
}

```