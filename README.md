# HTTP Request Logger API

This RESTful API, allows you to store & list a received HTTP requests to your application.

## Installation:

You need to install [MongoDB](https://www.digitalocean.com/community/tutorials/how-to-install-mongodb-on-ubuntu-18-04) and [GoLang](https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-18-04). 
After completing the necessary setups, copy `.env.example` to `.env` file and change these environment variable.

Then run;
```bash
go run main.go
```

## Authentication:

HTTP basic authentication with the credentials in your `.env` file.

## Create Request Log:

Request URI: **"/"** 
Request Type: **POST**

| Param | Type | Example|
| - | - | - |
| user_id | String (Required) | "424" |
| ip_address | String (Required) | "127.0.0.1" |
| domain | String (Required) | "https://example.com" |
| uri | String (Required) | "/products" |

## List All Requests:

Request URI: **"/list"**
Request Type: **GET**

| Param | Type | Example |
| - | - | - |
| domain | String (Required) | "String" |
| user | String | "424" |
| uri | String | "/products" |
| ip | String | "127.0.0.1" |
| date | String | "2020-10-13" |
| per-page | String | "10" |
| page | String | "1" |
| sort | String | "ASC" |


### Example:

```http request
http://127.0.0.1:8080/list?domain=https://example.com&user=424&ip=127.0.0.1&date=2020-10-19&per-page=5&page=1&sort=ASC
```

**Output**
```json
    "success": true,
    "data": [
        {
            "Id": "5f8d45a7546e959cecb3bbb8",
            "UserId": "424",
            "IpAddress": "127.0.0.1",
            "Uri": "/products",
            "Domain": "https://example.com",
            "CreatedAt": "2020-10-19T07:52:07.42Z"
        },
        {
            "Id": "5f8d45a9546e959cecb3bbb9",
            "UserId": "424",
            "IpAddress": "127.0.0.1",
            "Uri": "/products",
            "Domain": "https://example.com",
            "CreatedAt": "2020-10-19T07:52:09.075Z"
        },
        {
            "Id": "5f8d45a9546e959cecb3bbba",
            "UserId": "424",
            "IpAddress": "127.0.0.1",
            "Uri": "/products",
            "Domain": "https://example.com",
            "CreatedAt": "2020-10-19T07:52:09.513Z"
        },
        {
            "Id": "5f8d45aa546e959cecb3bbbb",
            "UserId": "424",
            "IpAddress": "127.0.0.1",
            "Uri": "/products",
            "Domain": "https://example.com",
            "CreatedAt": "2020-10-19T07:52:10.006Z"
        },
        {
            "Id": "5f8d45ac546e959cecb3bbbc",
            "UserId": "424",
            "IpAddress": "127.0.0.1",
            "Uri": "/products",
            "Domain": "https://example.com",
            "CreatedAt": "2020-10-19T07:52:12.601Z"
        }
    ],
    "extra": {
        "currentRows": 5,
        "hasNextPage": true,
        "page": 1,
        "perPage": 5,
        "totalPage": 8,
        "totalRows": 39
    }
```


## Postman Export

You can use this [export](https://raw.githubusercontent.com/sametsahindogan/request-logger/master/docs/request_logger.postman_collection.json).

## License
MIT © [Samet Sahindogan](https://github.com/sametsahindogan/request-logger/blob/master/LICENSE)