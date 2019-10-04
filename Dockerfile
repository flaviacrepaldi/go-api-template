FROM golang:1.12.6

ADD . /go/src/app
WORKDIR /go/src/app
COPY . .

RUN go get github.com/pilu/fresh
RUN go get -d -v -t ./...
RUN go install -v ./...