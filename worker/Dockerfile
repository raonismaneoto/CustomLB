FROM golang:1.14

WORKDIR /go/src/github.com/raonismaneoto/CustomLB/worker

RUN echo 'export PATH=$PATH:/usr/local/bin' >> ~/.bashrc

COPY ./worker ./
COPY ./go.mod ./
COPY ./go.sum ./

WORKDIR /go/src/github.com/raonismaneoto/CustomLB/worker

RUN go get -v
RUN go install -v

RUN go build -o main

CMD ('/bin/sh')