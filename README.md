# Baseball and Softball Scoring Service

We want you to build a simple baseball and softball score keeping REST API that mobile and web clients can use to save information about games. In its first iteration the service should provide endpoints to allow consumers to store and retrieve basic information about games and game events (like pitches and balls).

## Directions
- Please complete the open GitHub issue(s). Please commit your changes at least 24 hours before your visit and confirm your submission with your point of contact or let them know if you have any questions.
- Make sure you get to a working state. We're going to build some new features on top of your existing solution when you come into the office.
- We think it should normally take about a couple of hours to complete this assignment, but you are free to dedicate as much time as you see fit.
- The prompt is language-agnostic; feel free to choose the language(s) and technologies you are most comfortable with.
- We'll be running your software on our machines, so please include any setup instructions needed to run your solution.

## Instructions to run and test

The service uses an in memory db. Used below two variants of in memory db. Its configured in app.env file defaulted to memdb.
  - map based structure
  - <a href="https://github.com/hashicorp/go-memdb"> go-memdb </a>
```
To run the application
 cd back-end-naresh8t7
 go run cmd/main.go
 
```

```
To run the tests
 cd back-end-naresh8t7
 go test -v ./...
 
```

### Create Game End Point
```
Request:
curl --location --request POST 'http://localhost:8085/games' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "6690cf59-79de-445c-b9f7-04b7f1ee7990",
    "start": "2018-10-10T22:00:00.000Z",
    "end": "2018-10-11T01:00:00.000Z",
    "arrive": "2018-10-10T21:30:00.000Z"
}
'
Response:
{"status":"Successly created game"}
```

### Update Game End Point
Request:
curl --location --request PUT 'http://localhost:8085/games/6690cf59-79de-445c-b9f7-04b7f1ee7990' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "6690cf59-79de-445c-b9f7-04b7f1ee7990",
    "start": "2018-10-10T23:00:00.000Z",
    "end": "2018-10-11T01:00:00.000Z",
    "arrive": "2018-10-10T21:30:00.000Z"
}
Response:
{"status":"Succesfully updated game"}
'

### Get Game End Point
```
Request:
curl --location --request GET 'http://localhost:8085/games/6690cf59-79de-445c-b9f7-04b7f1ee7990'
Response:
{"id":"6690cf59-79de-445c-b9f7-04b7f1ee7990","start":"2018-10-10T22:00:00Z","end":"2018-10-11T01:00:00Z","arrive":"2018-10-10T21:30:00Z"}
```

### List Games Endpoint
```
Request:
curl --location --request GET 'http://localhost:8085/games'
Response:
[{"id":"6690cf59-79de-445c-b9f7-04b7f1ee7991","start":"2023-07-30T02:42:37.563855-04:00","end":"2023-07-30T05:42:37.563859-04:00","arrive":"2023-07-30T01:42:37.563859-04:00"},{"id":"6690cf59-79de-445c-b9f7-04b7f1ee7992","start":"2023-07-30T05:42:37.563859-04:00","end":"2023-07-30T08:42:37.563859-04:00","arrive":"2023-07-30T04:42:37.563859-04:00"},{"id":"6690cf59-79de-445c-b9f7-04b7f1ee7993","start":"2023-07-30T08:42:37.56386-04:00","end":"2023-07-30T11:42:37.56386-04:00","arrive":"2023-07-30T07:42:37.56386-04:00"},{"id":"6690cf59-79de-445c-b9f7-04b7f1ee7994","start":"2023-07-30T11:42:37.563861-04:00","end":"2023-07-30T14:42:37.563861-04:00","arrive":"2023-07-30T10:42:37.563861-04:00"}]
```

### Delete Game Endpoint
```
Request:
curl --location --request DELETE 'http://localhost:8085/games/6690cf59-79de-445c-b9f7-04b7f1ee7991'
Response:
Game with ID 6690cf59-79de-445c-b9f7-04b7f1ee7991 has been deleted
```

### Create Scoring Event EndPoint
```
Request:
curl --location --request POST 'http://localhost:8085/scoringEvents' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "486585db-75f2-467d-a825-b37777c96529",
    "game_id": "6690cf59-79de-445c-b9f7-04b7f1ee7990",
    "timestamp": "2018-10-10T22:03:56.413Z",
    "data": {
        "code": "pitch",
        "attributes": {
            "advances_count": true,
            "result": "ball_in_play"
        }
    }
}
'
Response:
{"status":"Successly created scoring event"}
```

### Update Scoring Event EndPoint
Request:
curl --location --request PUT 'http://localhost:8085/scoringEvents/486585db-75f2-467d-a825-b37777c96529' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "486585db-75f2-467d-a825-b37777c96529",
    "game_id": "6690cf59-79de-445c-b9f7-04b7f1ee7990",
    "timestamp": "2018-10-10T23:03:56.413Z",
    "data": {
        "code": "pitch",
        "attributes": {
            "advances_count": true,
            "result": "ball_in_play"
        }
    }
}
Response:
{"status":"Succesfully updated scoring event"}
'

### Get Scoring event End Point
```
Request:
curl --location --request GET 'http://localhost:8085/scoringEvents/486585db-75f2-467d-a825-b37777c96529'
Response:
{"id":"486585db-75f2-467d-a825-b37777c96529","game_id":"6690cf59-79de-445c-b9f7-04b7f1ee7990","timestamp":"2018-10-10T22:03:56.413Z","data":{"code":"pitch","attributes":{"advances_count":true,"result":"ball_in_play"}}}
```

### List Scoring Events Endpoint
```
Request:
curl --location --request GET 'http://localhost:8085/scoringEvents'
Response:
[{"id":"486585db-75f2-467d-a825-b37777c96530","game_id":"6690cf59-79de-445c-b9f7-04b7f1ee7991","timestamp":"2023-07-30T02:47:37.563861-04:00","data":{"code":"pitch","attributes":{"advances_count":true,"result":"ball_in_play"}}},{"id":"486585db-75f2-467d-a825-b37777c96531","game_id":"6690cf59-79de-445c-b9f7-04b7f1ee7992","timestamp":"2023-07-30T05:47:37.563861-04:00","data":{"code":"pitch","attributes":{"advances_count":true,"result":"ball_in_play"}}},{"id":"486585db-75f2-467d-a825-b37777c96532","game_id":"6690cf59-79de-445c-b9f7-04b7f1ee7993","timestamp":"2023-07-30T08:57:37.563861-04:00","data":{"code":"pitch","attributes":{"advances_count":true,"result":"ball_in_play"}}},{"id":"486585db-75f2-467d-a825-b37777c96533","game_id":"6690cf59-79de-445c-b9f7-04b7f1ee7994","timestamp":"2023-07-30T11:47:37.563862-04:00","data":{"code":"pitch","attributes":{"advances_count":true,"result":"ball_in_play"}}}]
```

### Delete Scoring Event Endpoint
```
Request:
curl --location --request DELETE 'http://localhost:8085/scoringEvents/486585db-75f2-467d-a825-b37777c96530'
Response:
Scoring event with ID 486585db-75f2-467d-a825-b37777c96530 has been deleted
```

### Metrics endpoint
```
Request:
curl --location --request GET 'http://localhost:8085/metrics'
Response:
Contains API metrics and various statistics related to go routines and gc etc.
```

### Health endpoint
```
Request:
curl --location --request GET 'http://localhost:8085/health'
Response:
<pre>Sat, 29 Jul 2023 23:45:14 -0400<br><br>Game Scoring Service. Status : OK<br></pre>
```

### Assumptions

The game id and scoring event id are unique. If the end points are called with same ids it will override the existing data as I am using in memory structure. But in real db application, we can throw an error.

### Libraries Used
- <a href="github.com/spf13/viper">Viper </a>: For configuring application params in config file.
- <a href="github.com/gorilla/mux"> Mux </a>: For http server implementation
- <a href="https://github.com/hashicorp/go-memdb"> go-memdB </a>: For In memory DB