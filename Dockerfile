FROM golang:1.8 as builder
WORKDIR /usr/src/app
COPY . .

RUN cd view && npm install && npm run build
RUN cd server && go get -d -v ./... && go build -o server .

EXPOSE 80
CMD ["./server/server"]

