FROM golang:1.14-alpine
WORKDIR /go/src/app
COPY . .
RUN mkdir build
RUN go build -o build/api-server cmd/server/main.go
RUN cp api-server ../../bin/
CMD ["/go/bin/api-server"]