FROM golang:1.18-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY .. .


RUN GOOS=linux CGO_ENABLED=0 go build -o /konntent-service-template ./cmd/app

FROM alpine
COPY --from=build /konntent-service-template ./service
COPY --from=build /app/.env ./app/.env

EXPOSE 8080

CMD [ "./service" ]