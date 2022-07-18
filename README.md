# urlservice

This is a Go service that counts word frequencies given a website target.

For fun, see output/tolstoy.txt, which is the result of `curl -X POST http://localhost:8081/url -d '{"target": "https://www.gutenberg.org/files/2600/2600-h/2600-h.htm"}' | jq . > tolstoy.txt`

# Build on docker

Run `make build` 

# Running on docker

Run `make run` after building the container (step above)

# Running tests

Run `make test`. Tests will also be run on build.

# Run from source

Run `go run main.go -dev` (you need to either specify -dev or configure TLS, otherwise the middleware will route to https on production mode (non -dev))

# Calling the API

There are two exposed endpoints:

`GET /health` 

Returns a placeholder response for server liveness verification

Example:

```
 curl -v http://localhost:8081/health

< HTTP/1.1 200 OK
< Referrer-Policy: no-referrer
< X-Content-Type-Options: nosniff
< X-Frame-Options: DENY
< X-Xss-Protection: 1; mode=block
< Date: Mon, 18 Jul 2022 02:49:20 GMT
< Content-Length: 0
<
* Connection #0 to host localhost left intact
```

`POST /url {"target": "<url>"}`

Returns a map of word frequency counts

Example:

```
curl -X POST http://localhost:8081/url -d '{"target": "https://example.com"}' | jq .

{
  "words": {
    "#38488f;": 1,
    "#f0f0f2;": 1,
    "#fdfdff;": 1,
    "(max-width": 1,
    "-apple-system": 1,
    "0": 1,
    "05em;": 1,
    "0;": 2,
    "2em;": 1,
    "2px": 2,
    "3px": 1,
    "5em": 1,
    "600px;": 1,
    "700px)": 1,
    "7px": 1,
    "@media": 1,
    "alink": 1,
    "arial": 1,
    "asking": 1,
    "auto;": 3,
    "avisited": 1,
    "background-color": 2,
    "blinkmacsystemfont": 1,
    "body": 1,
    "border-radius": 1,
    "box-shadow": 1,
    "color": 1,
    "coordination": 1,
    "div": 2,
    "documents": 1,
    "domain": 4,
    "example": 2,
    "examples": 1,
    "font-family": 1,
    "for": 2,
    "helvetica": 2,
    "html": 1,
    "illustrative": 1,
    "in": 3,
    "information": 1,
    "is": 1,
    "literature": 1,
    "margin": 3,
    "may": 1,
    "more": 1,
    "neue": 1,
    "none;": 1,
    "open": 1,
    "or": 1,
    "padding": 2,
    "permission": 1,
    "prior": 1,
    "rgba(000002);": 1,
    "sans": 1,
    "sans-serif;": 1,
    "segoe": 1,
    "system-ui": 1,
    "text-decoration": 1,
    "this": 2,
    "ui": 1,
    "use": 2,
    "width": 2,
    "without": 1,
    "you": 1,
    "{": 5,
    "}": 5
  }
}
```