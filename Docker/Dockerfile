FROM golang:1.19.3-alpine as build

WORKDIR /usr/local/go/src/userSL/

ADD https://github.com/vv198x/userSL/archive/refs/tags/v0.0.9.tar.gz .

RUN tar zxvf v0.0.9.tar.gz -C . --strip-components=1
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o userSL


FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/local/go/src/userSL/userSL /

ENTRYPOINT ["./userSL"]