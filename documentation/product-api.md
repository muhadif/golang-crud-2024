# Product - API Documentation

## Endpoint: Get Product Category

### Method: Get
- **URL**: `/product-category`
- **Body Format**: URL
- **Authentication**: Bearer


### Example Request
```
curl --request GET \
  --url http://localhost:8080/product-category \
  --header 'Authorization: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyU2VyaWFsIjoiVVNFUmY2NDRkYzc2Y2YiLCJSb2xlIjoidXNlciIsImV4cCI6MTcxNzE4ODI3N30.B_pxgSN5nsj0FyQYmWfl5RRdTBZ7Hg56YjjS7W10Fi4aF2saQ9cbqOZ3xqPgP-jcxhi4uhxEVyNWKf3ps42P7g'
```

### Response
#### Success (200)
```
{
  "status": "success",
  "data": [
    {
      "id": 1,
      "serial": "001",
      "name": "Electronics"
    }
  ]
}
```

## Endpoint: Get Product

### Method: Get
- **URL**: `/product`
- **Body Format**: URL
- **Authentication**: Bearer

### URL Query
- **productCategorySerial**: string, optional

### Example Request
```
curl --request GET \
  --url 'http://localhost:8080/product?productCategorySerial=002' \
  --header 'Authorization: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyU2VyaWFsIjoiVVNFUmY2NDRkYzc2Y2YiLCJSb2xlIjoidXNlciIsImV4cCI6MTcxNzE4ODI3N30.B_pxgSN5nsj0FyQYmWfl5RRdTBZ7Hg56YjjS7W10Fi4aF2saQ9cbqOZ3xqPgP-jcxhi4uhxEVyNWKf3ps42P7g'
  ```

### Response
#### Success (200)
```
{
  "status": "success",
  "data": [
    {
      "id": 2,
      "serial": "P002",
      "name": "T-shirt",
      "price": 19.99,
      "stock": 76,
      "description": "Comfortable cotton t-shirt in various colors.",
      "productCategories": [
        {
          "id": 2,
          "serial": "002",
          "name": "Clothing"
        }
      ]
    }
  ]
}
```