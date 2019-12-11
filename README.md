# RustAPI

## Introduction

### **warning this will run locally!!! for now**

## How to run

```console
usr/bandsAPI/go$ go run server.go
Server running on port 8080

usr/bandsAPI/go$
```

## REST

- URL: http://localhost:8080/api

The main resources are :

- http://localhost:8080/api/artists will show all the artists
- http://localhost:8080/api/locations will show all locations
- http://localhost:8080/api/dates will show all dates
- http://localhost:8080/api/relation will show relations between the locations and the dates

### Format example

- For artists.json

```json
{
    "index":[
        {
            "id": 1,
            "image" : "http://localhost:8080/api/images/queen.jpeg",
            "name" : "Queen",
            "members" : [
                "Freddie Mercury",
                "Brian May",
                "John Daecon",
                "Roger Meddows-Taylor",
                "Mike Grose",
                "Barry Mitchell",
                "Doug Fogie"
            ],
            "creationDate": 1970,
            "locations": "http://localhost:8080/api/locations/1",
            "consertDates": "http://localhost:8080/api/dates/1",
            "relations": "http://localhost:8080/api/relation/1"
        },
        {
            "id": 2,
            "image" : "http://localhost:8080/api/images/soja.jpeg",
            "name" : "SOJA",
            "members" : [
                "Jacob Hemphill",
                "Bob Jefferson",
                "Ryan \"Byrd\" Berty",
                "Ken Brownell",
                "Patrick O'Shea",
                "Hellman Escorcia",
                "Rafael Rodriguez",
                "Trevor Young"
            ],
            "creationDate": 1997,
            "locations": "http://localhost:8080/api/locations/2",
            "consertDates": "http://localhost:8080/api/dates/2",
            "relations": "http://localhost:8080/api/relation/2"
        }
    ]
}
```

- For locations.json

```json
{
    "index":[
        {
            "id": 1,
            "locations":[
                "north_carolina-usa",
                "georgia-usa",
                "los_angeles-usa",
                "saitama-japan",
                "osaka-japan",
                "nagoya-japan",
                "penrose-new_zealand",
                "dunedin-new_zealand"
            ],
            "dates":"http://localhost:8080/api/dates/1"
        },
        {
            "id": 2,
            "locations":[
                "playa_del_carmen-mexico",
                "papeete-french_polynesia",
                "noumea-new_caledonia"
            ],
            "dates":"http://localhost:8080/api/dates/2"
        }
    ]
}
```

- For dates.json

```json
{
    "index": [
        {
            "id":1,
            "dates":[
                "*23-08-2019",
                "*22-08-2019",
                "*20-08-2019",
                "*26-01-2020",
                "*28-01-2020",
                "*30-01-2019",
                "*07-02-2020",
                "*10-02-2020"
            ]
        },
        {
            "id":2,
            "dates":[
                "*05-12-2019",
                "06-12-2019",
                "07-12-2019",
                "08-12-2019",
                "09-12-2019",
                "*16-11-2019",
                "*15-11-2019"  
            ]
        }
    ]
}
```

## Possible queries

- `http://localhost:8080/api/artists/{id}`

`http://localhost:8080/api/artists/1`

```json
{
  "id": 1,
  "image": "http://localhost:8080/api/images/queen.jpeg",
  "name": "Queen",
  "members": [
    "Freddie Mercury",
    "Brian May",
    "John Daecon",
    "Roger Meddows-Taylor",
    "Mike Grose",
    "Barry Mitchell",
    "Doug Fogie"
  ],
  "creationDate": 1970,
  "locations": "http://localhost:8080/api/locations/1",
  "consertDates": "http://localhost:8080/api/dates/1",
  "relations": "http://localhost:8080/api/relation/1"
}
```

- `http://localhost:8080/api/locations/{id}`

`http://localhost:8080/api/locations/1`

```json
{
  "id": 1,
  "locations": [
    "north_carolina-usa",
    "georgia-usa",
    "los_angeles-usa",
    "saitama-japan",
    "osaka-japan",
    "nagoya-japan",
    "penrose-new_zealand",
    "dunedin-new_zealand"
  ]
}
```

- `http://localhost:8080/api/dates/{id}`

`http://localhost:8080/api/dates/1`

```json
{
  "id": 1,
  "dates": [
    "*23-08-2019",
    "*22-08-2019",
    "*20-08-2019",
    "*26-01-2020",
    "*28-01-2020",
    "*30-01-2019",
    "*07-02-2020",
    "*10-02-2020"
  ]
}
```

- `http://localhost:8080/api/relation/{id}`

`http://localhost:8080/api/relation/1`

```json
{
  "id": 1,
  "datesLocations": {
    "dunedin-new_zealand": [
      "10-02-2020"
    ],
    "georgia-usa": [
      "22-08-2019"
    ],
    "los_angeles-usa": [
      "20-08-2019"
    ],
    "nagoya-japan": [
      "30-01-2019"
    ],
    "north_carolina-usa": [
      "23-08-2019"
    ],
    "osaka-japan": [
      "28-01-2020"
    ],
    "penrose-new_zealand": [
      "07-02-2020"
    ],
    "saitama-japan": [
      "26-01-2020"
    ]
  }
}
```

- `http://localhost:8080/api/images/{name}`
