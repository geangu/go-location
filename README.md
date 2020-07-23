# Go-Location

This API saves a location in MongoDB and put this location in Redis cache,
also this API can returns the last location sent to Redis using an ID

## Run API

To run this API you need Docker and Docker Compose installed in your computer


```
docker-compose up --build
```

## Test end points

- Create new location
```
curl --location --request POST 'http://localhost:8080/location/123' \
--header 'Content-Type: application/json' \
--data-raw '{
    "longitude": "54.0214",
    "latitude": "30.0112"
}'

```
- Get last location
```
curl --location --request GET 'http://localhost:8080/location/123'
```

The last code was written as hobbie, if you need this code feel free to use ;)