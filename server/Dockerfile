FROM golang:1.12 as builder
WORKDIR /usr/src/app
COPY server/. .

RUN go get -d -v ./...
RUN go build -o server .

EXPOSE 80
CMD ["./server"]

