FROM golang:1.22-alpine AS build

RUN go version

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /ecom-go

FROM alpine:latest

WORKDIR /app

COPY --from=build /ecom-go /ecom-go

COPY .env .env

EXPOSE 8080

CMD [ "/ecom-go" ]