FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main main.go

EXPOSE 8081

CMD [ "./main" ]