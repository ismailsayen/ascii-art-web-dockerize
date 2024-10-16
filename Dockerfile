# to Downald the latest version of golang-- alpine is a version of linux with a small size 
FROM golang:alpine

WORKDIR /app

COPY . .

# this command is for install the terminal on our image
RUN apk add  bash

RUN go build -o main main.go

EXPOSE 8081

CMD [ "./main" ]
