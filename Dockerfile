## Compilation stage
FROM golang:1.12.1-alpine3.9
ENV GOPATH /go:/app
WORKDIR /app
ADD . /app
# get dependencies
RUN apk add git
RUN go get github.com/ollykel/webapp
RUN go get github.com/go-sql-driver/mysql
RUN go get gopkg.in/yaml.v2
# compile app
RUN go install serve
# IMPORTANT: src/client/public/bundle.js must be compiled before creating
# the Docker container.
# mv client's "public" dir to start of app dir
RUN mv src/client/public ./
# clean up image
RUN rm -rf ./src

## Final build stage
FROM alpine:3.9
WORKDIR /
COPY --from=0 /app ./app
WORKDIR /app
CMD ["bin/serve"]

