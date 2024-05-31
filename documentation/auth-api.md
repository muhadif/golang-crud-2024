# Auth -API Documentation

## Endpoint: Register

### Method: POST
- **URL**: `/auth/register`
- **Body Format**: JSON
- **Authentication**: None

### Request Body

```
{
  "email": "customer@gmail.com",
  "fullname": "customer",
  "username": "customer",
  "password": "12345678"
}
```

### Example Request
```
curl --request POST \
  --url http://localhost:8080/auth/register \
  --data '{
  "email": "customer@gmail.com",
  "fullname": "customer",
  "username": "customer",
  "password": "12345678"
}'
```

### Response
#### Success (200)
```
{
  "status": "success"
}
```

#### Email Taken (400)
```
{
  "status": "error",
  "error": "Email taken",
  "errorCode": 400
}
```

#### Validation Error (400)
```
{
  "status": "error",
  "error": "Key: 'RegisterRequest.FullName' Error:Field validation for 'FullName' failed on the 'required' tag",
  "errorCode": 400
}
```


## Endpoint: Login

### Method: POST
- **URL**: `/auth/login`
- **Body Format**: JSON
- **Authentication**: None

### Request Body

```
{
"email": "customer@gmail.com",
"password": "12345678"
}
```

### Example Request
```
curl -X POST http://localhost:8080/auth/login \
-H "Content-Type: application/json" \
-d '{
"email": "customer@gmail.com",
"password": "12345678"
}'
```

### Response
#### Success (200)
```
{
"status": "success",
"data": {
"access_token": {
"accessToken": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyU2VyaWFsIjoiVVNFUmY2NDRkYzc2Y2YiLCJSb2xlIjoidXNlciIsImV4cCI6MTcxNzE4ODI3N30.B_pxgSN5nsj0FyQYmWfl5RRdTBZ7Hg56YjjS7W10Fi4aF2saQ9cbqOZ3xqPgP-jcxhi4uhxEVyNWKf3ps42P7g",
"refreshToken": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.e30.BVLX2LHdmIRGgmxmWjC6gB_EJBhGPZ8ob6y8TN43DCuTtE3snkAW2Bp0BzVmh7EOE3dZWTK_Lt1P3zSotU2apQ",
"atExpires": 1717188277,
"rtExpires": 1717188277
}
}
}
```

#### Failed Login (412)
```
{
"status": "error",
"error": "Incorrect Login",
"errorCode": 412
}```


