FROM golang:latest

MAINTAINER Maintainer

ENV GIN_MODE=release
ENV PORT=3004

WORKDIR /go/src/go-docker-dev.to

COPY app /go/src/go-docker-dev.to/src

# Run the two commands below to install git and dependencies for the project.
# RUN apk update && apk add --no-cache git
# RUN go get ./...

COPY dependencies /go/src

RUN go build go-docker-dev.to/src/app

EXPOSE $PORT

ENTRYPOINT ["./app"]