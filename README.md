# paymentapi
This project is used for transaction functionality 
To start app you should install mysql in your database and create table in your enviornment

CREATE DATABASE IN IT (soubhagya)
=================================

create table like this
======================

user:

| userid | balance | status |
| ------ | ------- | ------ |
| 123    | 100     | 1      |
| 234    | 200     | 1      |

merchant:

| merchantid | balance | status |
| ---------- | ------- | ------ |
| 99987      | 91      | 1      |
| 6765       | 1900    | 1      |


user_transaction_status:

| txnid | userid | amount | balance | status | type |
|-------|--------|--------|---------|--------|-------|
|3567455 | 123   | 10     | 90      | SUCCESS  | DEBIT
|
|

merchant_transaction_status:

| txnid | userid | amount | balance | status | type |
|-------|--------|--------|---------|--------|-------|
|3567455 | 99987  | 9     |  10     | SUCCESS| CREDIT
================================


Should have go install of with GOPATH set

Run

go run main.go


===============
Sample Run ::



//user can send amount to merchant

1) 

POST request to "localhost:8000/send"

Request :

{
    "userid":123,
    "merchantid":99987,
    "amount":50
}
Response :

{
    "txnid": 2049586350,
    "status": "SUCCESS",
    "status_desc": "Successfully Credited in Merchant"
}

//merchant can send amount to user

2)POST :  localhost:8000/refund

Request :

{
    "userid":123,
    "merchantid":99987,
    "amount":50
}

Response :

{
    "txnid": 1974468106,
    "status": "SUCCESS",
    "status_desc": "Successfully Refunded"
}

//merchant can withdraw amount from his account(merchantid)

3) POST : localhost:8000/merhant-withraw

Request :
{
    "merchantid":99987,
    "amount":9
}
Response:

{
    "txnid": 9313955214,
    "status": "SUCCESS",
    "status_desc": "Withraw SuccessfullY"
}


//merchant can all transaction related to his account(merchantid)
4) POST : localhost:8000/merhant-txn-history

Request :
{
    "merchantid":99987
}
Response :
[
    {
        "txnid": 3082153551,
        "mechantid": 99987,
        "amount": 50,
        "balance": 100,
        "status": "SUCCESS",
        "type": "DEBIT"
    },
    {
        "txnid": 9313955214,
        "mechantid": 99987,
        "amount": 9,
        "balance": 91,
        "status": "SUCCESS",
        "type": "MERCHANT_WITHRAW"
    }
]




