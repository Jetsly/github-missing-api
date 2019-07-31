<h1 align="center">Github Missing Trending API</h1>

<p align="center">:octocat: A simple API  of Github.</p>


[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/huchenme/github-trending-api/blob/master/LICENSE)

## docker
```bash
# $ CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o github-missing-api .
$ docker build -t github-missing-api .
```

### API

#### List Languages

Receive popular languages and all languages.

**URL Endpoint:**

/trending/languages

**Response:**

```json
[
  {
    "id": "1c-enterprise",
    "name": "1C Enterprise"
  },
  {
    "id": "abap",
    "name": "ABAP"
  },
  {
    "id": "abnf",
    "name": "ABNF"
  },
  {
    "id": "actionscript",
    "name": "ActionScript"
  }
]
```


#### Trending Repositories

Receive an array of trending repositories.

**URL Endpoint:**

/trending/repositories?language=go&since=weekly

**Parameters:**

- `language`: **optional**, list trending repositories of certain programming languages, possible values are listed [here](languages.json).
- `since`: **optional**, default to `daily`, possible values: `daily`, `weekly` and `monthly`.

**Response:**

```json
[
  ...
  {
    "author": "google",
    "name": "gvisor",
    "avatar": "https://github.com/google.png",
    "url": "https://github.com/google/gvisor",
    "description": "Container Runtime Sandbox",
    "language": "Go",
    "languageColor": "#3572A5",
    "stars": 3320,
    "forks": 118,
    "currentPeriodStars": 1624,
    "builtBy": [
      {
        "href": "https://github.com/viatsko",
        "avatar": "https://avatars0.githubusercontent.com/u/376065",
        "username": "viatsko"
      }
    ]
  }
  ...
]
```



## contributors

* [github-trending-api](https://github.com/huchenme/github-trending-api)

