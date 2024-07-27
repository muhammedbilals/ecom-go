FROM golang:1.22-alpine AS build

RUN go version

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /ecom-go

FROM alpine:latest

COPY --from=build /ecom-go /ecom-go


EXPOSE 8080

CMD [ "/ecom-go" ]
