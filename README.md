### Available Games

#### 20 questions

Simple fun game where you should guess a noun by asking a maximum of 20 questions. Question can be answered only with "Yes", "No" and "Don't know".

### TODO

- UT
- 20 questions game. Word "crocodile".
- Choosing a language and supporting Google Translate API.

### How to Contribute

* Fork a repository
* Add/Fix something
* Check that tests are passing: `go test -v ./pkg/...`
* Create PR

### Run

```
docker pull pltvs/games
docker run -p 8080:8080 pltvs/games
// Navigate to http://localhost:8080
```

### Web Speech API

Web page uses [Web Speech API](https://dvcs.w3.org/hg/speech-api/raw-file/tip/speechapi.html).
