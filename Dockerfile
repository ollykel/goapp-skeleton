FROM golang:1.12.1-alpine3.9
ENV GOPATH /go:/app

WORKDIR /app

ADD . /app

# get dependencies
RUN apk add git
RUN go get github.com/ollykel/webapp
RUN go get github.com/go-sql-driver/mysql

RUN go install serve

# IMPORTANT: src/client/public/bundle.js must be compiled before creating
# the Docker container.

# mv client's "public" dir to start of app dir
RUN mv src/client/public /app

# clean up image
RUN rm -rf ./src
RUN rm -rf /go

CMD ["bin/serve"]

