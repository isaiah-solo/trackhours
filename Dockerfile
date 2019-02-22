FROM golang:1.8 as builder
WORKDIR /usr/src/app
COPY . .

RUN apt-get update && apt-get install -y --no-install-recommends apt-utils
RUN apt-get update
RUN apt-get -y install nodejs
RUN cd view && npm install && npm run build
RUN cd server && go get -d -v ./... && go build -o server .

EXPOSE 80
CMD ["./server/server"]

