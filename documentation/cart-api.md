# Cart and Checkout - API Documentation

## Endpoint: Create or Add Cart

### Method: POST
- **URL**: `/cart`
- **Body Format**: JSON
- **Authentication**: Bearer

### Request Body
| Field Name | Type   | Description             | Required |
|-----------|--------|-------------------------|----------|
| productSerial | string | serial from the product | Yes      |
| quantity | number | number of request       | Yes      |


```
{
  "productSerial" : "P002",
  "quantity": 2
}
```

### Example Request
```
curl --request POST \
  --url http://localhost:8080/cart \
  --header 'Authorization: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyU2VyaWFsIjoiVVNFUmY2NDRkYzc2Y2YiLCJSb2xlIjoidXNlciIsImV4cCI6MTcxNzE4ODI3N30.B_pxgSN5nsj0FyQYmWfl5RRdTBZ7Hg56YjjS7W10Fi4aF2saQ9cbqOZ3xqPgP-jcxhi4uhxEVyNWKf3ps42P7g' \
  --data '{
  "productSerial" : "P002",
  "quantity": 2
}'
```

### Response
#### Success (200)
```
{
  "status": "success"
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

#### Validation Error (400)
```
{
  "status": "error",
  "error": "InternalServerError",
  "errorCode": 500,
  "errorDescription": [
    {
      "Key": "InternalServerError",
      "Message": "Key: 'CreateCart.ProductSerial' Error:Field validation for 'ProductSerial' failed on the 'required' tag"
    }
  ]
}
```

## Endpoint: Get Cart

### Method: GET
- **URL**: `/cart`
- **Body Format**: -
- **Authentication**: Bearer


### Example Request
```
curl --request GET \
  --url http://localhost:8080/cart \
  --header 'Authorization: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyU2VyaWFsIjoiVVNFUmY2NDRkYzc2Y2YiLCJSb2xlIjoidXNlciIsImV4cCI6MTcxNzE4ODI3N30.B_pxgSN5nsj0FyQYmWfl5RRdTBZ7Hg56YjjS7W10Fi4aF2saQ9cbqOZ3xqPgP-jcxhi4uhxEVyNWKf3ps42P7g'
```

### Response
#### Success (200)
```
{
  "status": "success",
  "data": [
    {
      "id": 4,
      "userSerial": "USERf644dc76cf",
      "productSerial": "P002",
      "product": {
        "id": 2,
        "serial": "P002",
        "name": "T-shirt",
        "price": 19.99,
        "stock": 66,
        "description": "Comfortable cotton t-shirt in various colors."
      },
      "quantity": 21
    }
  ]
}
```

## Endpoint: Get Cart By ID

### Method: GET
- **URL**: `/cart/:id`
- **Body Format**: -
- **Authentication**: Bearer


### Example Request
```
curl --request GET \
  --url http://localhost:8080/cart/4 \
  --header 'Authorization: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyU2VyaWFsIjoiVVNFUmY2NDRkYzc2Y2YiLCJSb2xlIjoidXNlciIsImV4cCI6MTcxNzE4ODI3N30.B_pxgSN5nsj0FyQYmWfl5RRdTBZ7Hg56YjjS7W10Fi4aF2saQ9cbqOZ3xqPgP-jcxhi4uhxEVyNWKf3ps42P7g'
```

### Response
#### Success (200)
```
{
  "status": "success",
  "data": {
    "id": 4,
    "userSerial": "USERf644dc76cf",
    "productSerial": "P002",
    "product": {
      "id": 2,
      "serial": "P002",
      "name": "T-shirt",
      "price": 19.99,
      "stock": 66,
      "description": "Comfortable cotton t-shirt in various colors."
    },
    "quantity": 1000
  }
}
```

## Endpoint: Update Cart

### Method: PUT
- **URL**: `/cart`
- **Body Format**: JSON
- **Authentication**: Bearer

### Request Body

```
{
  "id": 4,
  "productSerial" : "P002",
  "quantity": 10
}
```

### Example Request
```
curl --request PUT \
  --url http://localhost:8080/cart \
  --header 'Authorization: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyU2VyaWFsIjoiVVNFUmY2NDRkYzc2Y2YiLCJSb2xlIjoidXNlciIsImV4cCI6MTcxNzE4ODI3N30.B_pxgSN5nsj0FyQYmWfl5RRdTBZ7Hg56YjjS7W10Fi4aF2saQ9cbqOZ3xqPgP-jcxhi4uhxEVyNWKf3ps42P7g' \
  --data '{
  "id": 2,
  "productSerial" : "P004",
  "quantity": 0
}'
```

### Response
#### Success (200)
```
{
  "status": "success"
}
```

#### Cart Not Found (404)
```
{
  "status": "error",
  "error": "cart not found",
  "errorCode": 404
}
```

### Product out of stock (400)
```
{
  "status": "error",
  "error": "product stock lower than request",
  "errorCode": 400
}
```

## Endpoint: Delete Cart

### Method: DELETE
- **URL**: `/cart`
- **Body Format**: JSON
- **Authentication**: Bearer

### Request Body

```
{
  "id": 3,
  "productSerial" : "P004"
}
```

### Example Request
```
curl --request DELETE \
  --url http://localhost:8080/cart \
  --header 'Authorization: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyU2VyaWFsIjoiVVNFUmY2NDRkYzc2Y2YiLCJSb2xlIjoidXNlciIsImV4cCI6MTcxNzE4ODI3N30.B_pxgSN5nsj0FyQYmWfl5RRdTBZ7Hg56YjjS7W10Fi4aF2saQ9cbqOZ3xqPgP-jcxhi4uhxEVyNWKf3ps42P7g' \
  --data '{
  "id": 1,
  "productSerial" : "P004"
}'
```

### Response
#### Success (200)
```
{
  "status": "success"
}
```

#### Cart Not Found (404)
```
{
  "status": "error",
  "error": "cart not found",
  "errorCode": 404
}
```

## Endpoint: Create Checkout Session

### Method: POST
- **URL**: `/checkout`
- **Body Format**: JSON
- **Authentication**: Bearer

### Request Body

| Field Name | Type   | Description                        | Required |
|-----------|--------|------------------------------------|----------|
| cartId | string | id result from add product to cart | Yes      |

```
{
  "cartItems" : [
    {
      "cartId": 4
    }
  ]
}
```

### Example Request
```
curl --request POST \
  --url http://localhost:8080/checkout \
  --header 'Authorization: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyU2VyaWFsIjoiVVNFUmY2NDRkYzc2Y2YiLCJSb2xlIjoidXNlciIsImV4cCI6MTcxNzE4ODI3N30.B_pxgSN5nsj0FyQYmWfl5RRdTBZ7Hg56YjjS7W10Fi4aF2saQ9cbqOZ3xqPgP-jcxhi4uhxEVyNWKf3ps42P7g' \
  --data '{
  "cartItems" : [
    {
      "cartId": 4
    }
  ]
}'
```

### Response
#### Success (200)
```
{
  "status": "success"
}
```

#### Cart Not Found (404)
```
{
  "status": "error",
  "error": "cart not found",
  "errorCode": 404
}
```

## Endpoint: Get Current Checkout Session

### Method: GET
- **URL**: `/checkout`
- **Body Format**: JSON
- **Authentication**: Bearer

### Example Request
```
curl --request GET \
  --url http://localhost:8080/checkout \
  --header 'Authorization: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyU2VyaWFsIjoiVVNFUmY2NDRkYzc2Y2YiLCJSb2xlIjoidXNlciIsImV4cCI6MTcxNzE4ODI3N30.B_pxgSN5nsj0FyQYmWfl5RRdTBZ7Hg56YjjS7W10Fi4aF2saQ9cbqOZ3xqPgP-jcxhi4uhxEVyNWKf3ps42P7g'
```

### Response
#### Success (200)
```
{
  "status": "success",
  "data": {
    "cartItems": [
      {
        "id": 4,
        "userSerial": "USERf644dc76cf",
        "productSerial": "P002",
        "product": {
          "id": 2,
          "serial": "P002",
          "name": "T-shirt",
          "price": 19.99,
          "stock": 66,
          "description": "Comfortable cotton t-shirt in various colors."
        },
        "quantity": 1000
      }
    ],
    "total": 19990
  }
}
```