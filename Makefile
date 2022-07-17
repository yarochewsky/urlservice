IMAGE := "url_service"

clean:
	rm -rf target

.PHONY: compile
compile:
	GOARCH=amd64 GOOS=linux go build -o target/main.linux main.go

deploy: build docker-build
	@echo "pushed"

build: compile
	docker build --no-cache -t ${IMAGE} .

run:
	docker run -t ${IMAGE}
