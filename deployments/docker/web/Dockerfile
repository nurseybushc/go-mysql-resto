# Telling to use Docker's golang ready image
FROM golang:alpine
#EXPOSE 8080

#RUN apt-get update && apt-get install -y mysql-client
RUN apk update && apk add mysql git
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/nurseybushc/go-mysql-resto
