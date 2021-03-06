<h1 align="center">Github Missing API</h1>

<p align="center">:octocat: A simple API  of Github.</p>


[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/huchenme/github-trending-api/blob/master/LICENSE)

## docker
```bash
$ docker run --rm -it -p 8080:8000 ddotjs/github-missing-api 
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

- `language`: **optional**, list trending repositories of certain programming languages.
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

### Trending Developers

Receive an array of trending developers.

**URL Endpoint:**

/trending/developers?language=javascript&since=weekly

**Parameters:**

- `language`: **optional**, list trending repositories of certain programming languages.
- `since`: **optional**, default to `daily`, possible values: `daily`, `weekly` and `monthly`.

**Response:**

```json
[
  {
    "username": "google",
    "name": "Google",
    "type": "organization",
    "url": "https://github.com/google",
    "avatar": "https://avatars0.githubusercontent.com/u/1342004",
    "repo": {
      "name": "traceur-compiler",
      "description": "Traceur is a JavaScript.next-to-JavaScript-of-today compiler",
      "url": "https://github.com/google/traceur-compiler"
    }
  }
]
```

> `type` could be `organization` or `user`.


## contributors

* [github-trending-api](https://github.com/huchenme/github-trending-api)

