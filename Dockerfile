FROM golang:1.17-alpine3.13 as builder
WORKDIR /code
COPY go.mod go.sum /code/

RUN go mod download
COPY . /code
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/code code

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder code/build/code /usr/bin/code
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/code"]