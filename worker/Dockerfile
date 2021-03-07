FROM golang:1.14

WORKDIR /go/src/github.com/raonismaneoto/CustomLB/

RUN echo 'export PATH=$PATH:/usr/local/bin' >> ~/.bashrc

COPY ./ ./
COPY ../go.mod ./
COPY ../go.sum ./

WORKDIR /go/src/github.com/raonismaneoto/CustomLB/

RUN go get -v
RUN go install -v

RUN go build -o main

CMD ('/bin/sh')