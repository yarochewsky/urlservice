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