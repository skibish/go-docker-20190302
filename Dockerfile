FROM golang:1.11.5 as build

WORKDIR /go/src/github.com/skibish/go-docker-20190302

COPY . .

RUN cd cmd/chatserver && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/chatserver

FROM scratch

COPY --from=build /go/bin/chatserver /chatserver

CMD ["/chatserver"]
