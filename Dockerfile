FROM golang:alpine

#Author
MAINTAINER EmanueleGallone

ENV GIN_MODE=release
ENV PORT=80
ENV MYSQL_USER="root"
ENV MYSQL_PASSWORD="root"
ENV MYSQL_HOST="172.17.0.2"
ENV MYSQL_DB="users_db"

# Create working folder
RUN mkdir /app
COPY . /app

RUN apk update && apk add git

WORKDIR /app

RUN go build -o main

EXPOSE $PORT

CMD ["/app/main"]