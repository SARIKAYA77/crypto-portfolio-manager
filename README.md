### Requirements
* Go Language
* Postgres or MongoDB
* Docker
* Git
* Go Module
* [CoinMarketCap API](https://coinmarketcap.com/api/)

You need to develop a REST API and integrate 3rd part API which is used to manage a cryptocurrency portfolio service with CoinMarketCap. It has some basic CRUD operations on currency system.

Let's say, it's listening on localhost at port 8080;

These are example requests and responses it needs to provide;

---

#### Add a currency
Create a new portfolio data with given parameters and return the total price of your currency.
##### Sample Request
```
curl -X PUT \
 -d '{"code": "BTC", "amount": "5"}' \
 -H 'Content-Type: application/json' \
  http://localhost:8080/currency
```

##### Errors:

| Status Code | Description | Sample Response |
| --  | -- | -- |
| 200 | Success | {"id": 1, "code": "BTC", "amount": 5, price: "302877.0"} |
| 400 | When request body or parameters wrong | {"error": "Bad request"}|
| 403 | If currency already exists | {"error": "Currency already exists"} |
| 500 | When something unexpected happens | {"error": "server error"} |

---

#### Edit a currency's attributes
Edit existing portfolio data with given currency id
##### Sample Request
```
curl -X PATCH \
 -d '{"code": "BTC", "amount": "4"}' \
 -H 'Content-Type: application/json' \
  http://localhost:8080/currencies/1
```

##### Errors:

| Status Code | Description | Sample Response |
| --  | -- | -- |
| 200 | Success | {"id": 1, "code": "BTC", "amount": "4", price: "242301.6"} |
| 400 | When request body or parameters wrong | {"error": "Bad request"}|
| 404 | If currency not found | {"error": "Currency with that id does not exist"} |
| 500 | When something unexpected happens | {"error": "server error"} |

---

#### Delete a currency
Delete exist portflio data with given id
##### Sample Request
```
curl -X DELETE \
  http://localhost:8080/currencies/1
```

##### Errors:

| Status Code | Description | Sample Response |
| --  | -- | -- |
| 200 | Success |  |
| 400 | When request body or parameters wrong | {"error": "Bad request"}|
| 404 | If currency not found | {"error": "Currency with that id does not exist"} |
| 500 | When something unexpected happens | {"error": "server error"} |

---

#### Find a currency with ID
Get existing portfolio data with historical price changes.

##### Sample Request
```
curl -X GET \
  http://localhost:8080/currencies/1
```

##### Errors:

| Status Code | Description | Sample Response |
| --  | -- | -- |
| 200 | Success | {"id": 1,"code": "BTC","history": [{"amount": 4,"price": {"old": 242301.6,"current": 255064.0}},{"amount":5,"price": {"old": 302877.0,"current": 315353.3}}]} |
| 400 | When request body or parameters wrong | {"error": "Bad request"}|
| 404 | If currency not found | {"error": "Currency with that id does not exist"} |
| 500 | When something unexpected happens | {"error": "server error"} |


---

#### Get All Currencies

##### Sample Request
```
curl -X GET \
  http://localhost:8080/currencies
```

##### Errors:

| Status Code | Description | Sample Response |
| --  | -- | -- |
| 200 | Success | [{"id": 1,"code": "BTC","history": [{"amount": 4,"price": {"old": 242301.6,"new": 255064.0}},{"amount":5,"price": {"old": 302877.0,"current": 315353.3}}]},{"id": 2,"code": "ETH","history": [{"amount": 10,"price": {"old": 42610.0,"current": 4361.4}},{"amount":2,"price": {"old": 8522.0,"current": 8425.0}}]}] |
| 400 | When request body or parameters wrong | {"error": "Bad request"}|
| 404 | If currency not found | {"error": "Currency with that id does not exist"} |
| 500 | When something unexpected happens | {"error": "server error"} |


--- 

### Expectations

* This service should be written in Go lang, you can use any framework or library you need.
* There should be a database system to store currency data, preferably Postgres or MongoDB.
* Dockerfile should be provided in the project.
* Code Quality and Design

#### Extras (Optional)

* Kubernetes deployment setup (yaml file[s])
* Unit tests
