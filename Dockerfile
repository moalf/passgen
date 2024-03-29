FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /passgen

EXPOSE 8080

RUN chmod +x /passgen

CMD [ "/passgen" ]