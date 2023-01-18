Wallet Balance API
This API is a simple application for inserting, updating, and getting the balance of your wallet.

Getting Started
To get started with the API, you will need to clone the repository and install the necessary dependencies.

Prerequisites
Go 1.14 or later
MySQL
Installing
1. Clone the repository
git clone https://github.com/nurcholisnanda/wallet-record.git
2. Install the dependencies
go mod download or go mod tidy
3. Create the MySQL database
4. Update the MySQL credentials in the .env file
5. Start the server
go run main.go

API Endpoints
Visit https://nimble-monument-374407.an.r.appspot.com/swagger/index.html to open swagger doc
The API has the following endpoints:

Insert Record
POST /records

Parameters
Name	Type	Description
datetime	timestamp	Timestamp of the transaction
amount	float	Amount of the transaction

Example Request
POST /records
{
	"datetime":"2022-12-10T14:00:00+07:00",
	"amount": 10000
}
Example Response
{
    "success": true,
    "status": 201,
    "message": "success"
}

Get Latest Record
GET /records/latest
Example Response
{
    "datetime": "2022-12-10T09:00:00Z",
    "amount": 10000
}

Get History
POST /records/history
Parameters
Name	Type	Description
startDatetime	timestamp	Start timestamp of the history
endDatetime	timestamp	End timestamp of the history
POST /records/history
{
	"startDatetime":"2022-12-10T08:00:00+00:00",
	"endDatetime":"2022-12-10T09:00:00+00:00"
}
#### Example Response
[
{
"id": 1,
"datetime": "2022-12-10T09:00:00Z",
"amount": 10000
}
]


## Built With

- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [GORM](https://github.com/jinzhu/gorm) - ORM
- [Swaggo](https://github.com/swaggo/swag) - API documentation

## Authors

- **Nanda Nurcholis** - [nurcholisnanda](https://github.com/nurcholisnanda)


