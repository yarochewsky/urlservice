IMAGE := "url_service"

clean:
	rm -rf target

.PHONY: compile
compile: test
	GOARCH=amd64 GOOS=linux go build -o target/main.linux main.go

test:
	go test -race ./...

deploy: build docker-build
	@echo "pushed"

build: compile
	docker build --no-cache -t ${IMAGE} .

run:
	docker run -p 8081:8081 -t ${IMAGE}
