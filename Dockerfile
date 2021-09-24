FROM golang:1.17

WORKDIR /go/src/go-gin-rest-boilerplate
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go" , "run", "main.go"]
