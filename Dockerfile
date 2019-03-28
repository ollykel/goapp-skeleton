FROM golang:1.10
ENV GOPATH /go:/app

WORKDIR /app

ADD . /app

# get dependencies
RUN go get github.com/ollykel/webapp
RUN go get github.com/go-sql-driver/mysql

RUN go install serve

CMD ["bin/serve"]

