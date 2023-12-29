FROM golang:1.21-alpine

WORKDIR /opt/app/api

RUN go install github.com/cosmtrek/air@latest

ENTRYPOINT ["air", "--build.cmd", "go build -o ./.bin/sports_score ./main.go", "--build.bin", "./.bin/sports_score"]

EXPOSE 8080
