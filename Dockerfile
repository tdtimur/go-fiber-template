FROM golang:1.14-alpine as build-deps
WORKDIR /go/src/app
COPY . .
RUN mkdir build
RUN go build -o build/server main.go

FROM alpine:latest
COPY --from=build-deps /go/src/app/build/server /usr/local/bin/
CMD ["server"]