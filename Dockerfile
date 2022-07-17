FROM golang:11

WORKDIR /service

ADD target/main.linux /service/main

CMD ["main"]