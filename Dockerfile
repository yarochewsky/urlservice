FROM alpine

WORKDIR /service

ADD target/main.linux /service/main

CMD ["./main", "-dev"]