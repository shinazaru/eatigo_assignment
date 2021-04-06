# eatigo_assignment
## Probelm
```
Write a production ready microservice in Go, which will function as a link shortener service like Bitly or TinyURL. The service will need to be able to convert a provided URL into a fictional, shortened URL. No front end will be required, but the service needs to expose production ready APIs.

Some rules:
1. Each shortened URL will remain stored for 30 days.
    a. Bonus points for coming up with a refreshing caching solution, where shortened URLs are stored longer if they’ve been used/accessed within its 30 day limit.
    b. Keep in mind that URLs can be added more than once, or that shortened strings need to be unique.
2. The service does require authentication.
    a. Bonus points for OAuth.
3. You can choose any data storage you like, but it should be a good choice for the task given.
```
## Solution
```
    I store a URI with mongoDB and create an index of collection with {{ "expireAt: 1"}, {expireAfterSecond: 0}}
    for use a TTL setting, It's a mongo feature,
    When document is create, I will put current time with 30 days into "expireAt"
    After that, When expireAt is older than time.Now(). document will be deleted. 

    When generated URI is accessed, "expireAt" will be update with "current time with 30 days"
    That like an extend that link life time,

    ps. i decided to use mongo, cause i need to use an existing feature of tools,
    Why not Redis? : mongo will be easier to keep a data when service or some node in system is down, 
```

## Library Stack
- Echo
- mongo-driver
- shortuuid

## API Spec
### create short link
- METHOD: POST
- ENDPOINT: http(s)://${domain}/
- RequestBody
``` json
{
    "uri": "http://www.example.com"
}
```
- RespondsBody
``` json
{
    "uri": "http(s)://${domain}/${shotID}"
}
```

### call short link to actually link
- METHOD: GET
- ENDPOINT: http(s)://${domain}/${shotID}

### call short link to actually link
- METHOD: DELETE
- ENDPOINT: http(s)://${domain}/${shotID}
- RespondsBody
``` 
DELETE
```

## Environment Params List
- MONGO_URL: uri for db connection ex. mongodb://eatigo:ogitae@db:27017
- MONGO_DATABASE: "database name"
- API_DOMAIN: "api domain"