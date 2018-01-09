FROM golang:1.9

WORKDIR /go/src/crawler

ADD ./crawler .
VOLUME [ "/go/src/crawler/data" ]

RUN go build
CMD ./crawler
