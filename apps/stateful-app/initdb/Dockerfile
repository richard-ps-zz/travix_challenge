FROM golang:1.12-alpine
WORKDIR /go/src/app
COPY src .
COPY create_table.sql .
COPY articles.sql .

RUN apk add --no-cache git

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]