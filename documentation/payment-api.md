# Payment - API Documentation

## Endpoint: Create Payment from Current Checkout Session

### Method: POST
- **URL**: `/payment`
- **Body Format**: JSON
- **Authentication**: Bearer

### Request Body

```
{
  "paymentMethod" : "VA_TRANSFER"
}
```

### Example Request
```
curl --request POST \
  --url http://localhost:8080/payment \
  --header 'Authorization: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyU2VyaWFsIjoiVVNFUmY2NDRkYzc2Y2YiLCJSb2xlIjoidXNlciIsImV4cCI6MTcxNzE4ODI3N30.B_pxgSN5nsj0FyQYmWfl5RRdTBZ7Hg56YjjS7W10Fi4aF2saQ9cbqOZ3xqPgP-jcxhi4uhxEVyNWKf3ps42P7g' \
  --data '{
  "paymentMethod" : "VA_TRANSFER"
}'
```

### Response
#### Success (200)
```
{
  "status": "success",
  "data": {
    "serial": "PAY-d1f87",
    "openTime": "2024-05-31T13:30:41.198331+07:00",
    "expiredTime": null,
    "userSerial": "USERf644dc76cf",
    "totalPrice": 199.89999999999998,
    "PaymentItems": [
      {
        "paymentHistorySerial": "PAY-d1f87",
        "productSerial": "P002",
        "price": 19.99,
        "quantity": 10,
        "product": {
          "id": 2,
          "serial": "P002",
          "name": "T-shirt",
          "price": 19.99,
          "stock": 56,
          "description": "Comfortable cotton t-shirt in various colors."
        }
      }
    ],
    "paymentMethod": "VA_TRANSFER",
    "status": "WAITING",
    "transactionId": "a9de81978d41bb2fe44d83c606010b6d"
  }
}
```

#### Product out of Stock (400)
```
{
  "status": "error",
  "error": "product stock lower than request",
  "errorCode": 400
}
```
#### Product not exist (404)
```
{
  "status": "error",
  "error": "product not found",
  "errorCode": 404
}
```

#### Validation Error on Payment Method (400)
```
{
  "status": "error",
  "error": "Key: 'CreatePaymentRequest.PaymentMethod' Error:Field validation for 'PaymentMethod' failed on the 'enum' tag",
  "errorCode": 412
}
```

### Checkout is empty (412)
```
{
  "status": "error",
  "error": "need checkout first before making payment",
  "errorCode": 400
}
```

## Endpoint: Get Payment History

### Method: GET
- **URL**: `/payment/history`
- **Body Format**: JSON
- **Authentication**: Bearer


### Example Request
```
curl --request GET \
  --url http://localhost:8080/payment/history \
  --header 'Authorization: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyU2VyaWFsIjoiVVNFUmY2NDRkYzc2Y2YiLCJSb2xlIjoidXNlciIsImV4cCI6MTcxNzE4ODI3N30.B_pxgSN5nsj0FyQYmWfl5RRdTBZ7Hg56YjjS7W10Fi4aF2saQ9cbqOZ3xqPgP-jcxhi4uhxEVyNWKf3ps42P7g'
```

### Response
#### Success (200)
```
{
  "status": "success",
  "data": [
    {
      "serial": "PAY-9c78a",
      "openTime": "2024-05-31T15:06:47+07:00",
      "expiredTime": null,
      "userSerial": "USERf644dc76cf",
      "totalPrice": 39.98,
      "PaymentItems": [
        {
          "paymentHistorySerial": "PAY-9c78a",
          "productSerial": "P002",
          "price": 19.99,
          "quantity": 2,
          "product": {
            "id": 2,
            "serial": "P002",
            "name": "T-shirt",
            "price": 19.99,
            "stock": 64,
            "description": "Comfortable cotton t-shirt in various colors."
          }
        }
      ],
      "paymentMethod": "VA_TRANSFER",
      "status": "WAITING",
      "transactionId": "e38812a010cdb526b4c6f374d0be18fc"
    }
  ]
}
```

## Endpoint: Cancel Payment

### Method: GET
- **URL**: `/payment/cancel`
- **Body Format**: JSON
- **Authentication**: Bearer


### Request Body 
```
{
  "paymentSerial" : "PAY-d1f87"
}
```
### Example Request
```
curl --request POST \
  --url http://localhost:8080/payment/cancel \
  --header 'Authorization: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyU2VyaWFsIjoiVVNFUmY2NDRkYzc2Y2YiLCJSb2xlIjoidXNlciIsImV4cCI6MTcxNzE4ODI3N30.B_pxgSN5nsj0FyQYmWfl5RRdTBZ7Hg56YjjS7W10Fi4aF2saQ9cbqOZ3xqPgP-jcxhi4uhxEVyNWKf3ps42P7g' \
  --data '{
  "paymentSerial" : "PAY-d1f87"
}'
```

### Response
#### Success (200)
```
{
  "status": "success"
}
```

### Payment not found (404)
```
{
  "status": "error",
  "error": "payment not found",
  "errorCode": 404
}
```

## Endpoint: Payment Callback (VA-Transfer)

### Method: GET
- **URL**: `/callback/payment/va-transfer`
- **Body Format**: JSON
- **Authentication**: -

### Request Body
```
{
  "transactionId" : "1fbde43e9f3ec0d247cbe7bccbab6ca7",
  "status": "SUCCESS"
}
```
### Example Request
```
curl --request POST \
  --url http://localhost:8080/callback/payment/va-transfer \
  --header 'Authorization: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyU2VyaWFsIjoiVVNFUjZlZmVkNDVlZGEiLCJSb2xlIjoidXNlciIsImV4cCI6MTcxNzEwNTM3OH0.PW5RUZmnxpKm7cM8HaagQLBrtMCUotsiSzLUyuUBO48n200Wh0bjxueqqtWj_JlGxadUUam97K66bY09KJl8Zg' \
  --data '{
  "transactionId" : "1fbde43e9f3ec0d247cbe7bccbab6ca7",
  "status": "SUCCESS"
}'
```

### Response
#### Success (200)
```
{
  "status": "success"
}
```

### Payment not found (404)
```
{
  "status": "error",
  "error": "payment not found",
  "errorCode": 404
}
```
