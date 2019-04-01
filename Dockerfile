FROM golang:1.10
ENV GOPATH /go:/app

WORKDIR /app

ADD . /app

# get dependencies
RUN go get github.com/ollykel/webapp
RUN go get github.com/go-sql-driver/mysql

RUN go install serve

# IMPORTANT: src/client/public/bundle.js must be compiled before creating
# the Docker container.
# This is done to maintain a minimal disk image, without client source code.

# link to client's "public" dir
RUN ln -s src/client/public public

CMD ["bin/serve"]

