FROM golang:alpine

#Author
MAINTAINER EmanueleGallone

ENV GIN_MODE=release
ENV PORT=80

# Create working folder
RUN mkdir /app
COPY . /app

RUN apk update && apk add git

WORKDIR /app

RUN go build

EXPOSE $PORT

ENTRYPOINT ["go run main.go"]