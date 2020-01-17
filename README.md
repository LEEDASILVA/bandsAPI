# RestAPI

## Introduction

### [RESTApi](https://groupietrackers.herokuapp.com/api)

## How to run

```console
usr/bandsAPI/go$ go run server.go
Server running on port 8080

usr/bandsAPI/go$
```

## REST

- URL: `https://groupietrackers.herokuapp.com/api`

The main resources are :

- `https://groupietrackers.herokuapp.com/api/artists` will show all the artists
- `https://groupietrackers.herokuapp.com/api/locations` will show all locations
- `https://groupietrackers.herokuapp.com/api/dates` will show all dates
- `https://groupietrackers.herokuapp.com/api/relation` will show relations between the locations and the dates

### Format example

- For artists.json

```json
{
    "index":[
        {
            "id": 1,
            "image" : "https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
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
            "locations": "https://groupietrackers.herokuapp.com/api/locations/1",
            "consertDates": "https://groupietrackers.herokuapp.com/api/dates/1",
            "relations": "https://groupietrackers.herokuapp.com/api/relation/1"
        },
        {
            "id": 2,
            "image" : "https://groupietrackers.herokuapp.com/api/soja.jpeg",
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
            "locations": "https://groupietrackers.herokuapp.com/api/locations/2",
            "consertDates": "https://groupietrackers.herokuapp.com/api/dates/2",
            "relations": "https://groupietrackers.herokuapp.com/api/relation/2"
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
            "dates":"https://groupietrackers.herokuapp.com/api/dates/1"
        },
        {
            "id": 2,
            "locations":[
                "playa_del_carmen-mexico",
                "papeete-french_polynesia",
                "noumea-new_caledonia"
            ],
            "dates":"https://groupietrackers.herokuapp.com/api/dates/2"
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

- `https://groupietrackers.herokuapp.com/api/artists/{id}`

`https://groupietrackers.herokuapp.com/api/artists/1`

```json
{
  "id": 1,
  "image": "https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
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
  "locations": "https://groupietrackers.herokuapp.com/api/locations/1",
  "consertDates": "https://groupietrackers.herokuapp.com/api/dates/1",
  "relations": "https://groupietrackers.herokuapp.com/api/relation/1"
}
```

- `https://groupietrackers.herokuapp.com/api/locations/{id}`

`https://groupietrackers.herokuapp.com/api/locations/1`

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

- `https://groupietrackers.herokuapp.com/api/dates/{id}`

`https://groupietrackers.herokuapp.com/api/dates/1`

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

- `https://groupietrackers.herokuapp.com/api/relation/{id}`

`https://groupietrackers.herokuapp.com/api/relation/1`

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

- `https://groupietrackers.herokuapp.com/api/images/{name}`
