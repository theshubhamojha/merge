FROM golang:1.19-alpine

WORKDIR /app
COPY . .

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
    
RUN go mod tidy
RUN go build

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

EXPOSE 80
EXPOSE 443
EXPOSE 8080

CMD [ "go", "run", "main.go" ]
