FROM golang:1.12.7
MAINTAINER owner

WORKDIR /go/src/search
RUN go get -u github.com/golang/dep/cmd/dep

COPY ./Gopkg.lock .
COPY ./Gopkg.toml .

RUN dep ensure -v -vendor-only

COPY . .

WORKDIR go/src/search/cmd/search/

RUN go build

CMD ["./search"]
