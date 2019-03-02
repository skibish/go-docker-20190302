FROM golang:1.11.5

WORKDIR /go/src/github.com/skibish/go-docker-20190302

COPY . .

RUN cd cmd/chatserver && go install

CMD ["chatserver"]
