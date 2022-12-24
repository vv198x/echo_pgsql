FROM golang:1.19.3-alpine

WORKDIR /usr/local/go/src/userSL/

ADD https://github.com/vv198x/userSL/archive/refs/tags/v0.0.5.tar.gz .

RUN tar zxvf v0.0.5.tar.gz -C /usr/local/go/src/userSL/ --strip-components=1
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o userSL

RUN ./userSL init
ENTRYPOINT ["./userSL"]
