FROM golang:1.19.3-alpine

COPY build1 /userSL/
COPY table.sql /userSL/

WORKDIR /userSL

ENTRYPOINT ["./build1"]
