# eatigo_assignment

## API Spec
### create short link
- METHOD: POST
- ENDPOINT: http(s)://${domain}/
- Authentication: basic-auth
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