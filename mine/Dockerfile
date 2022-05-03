FROM golang:1.18-alpine as build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o /usr/local/bin/app ./...

FROM alpine

WORKDIR /

COPY --from=build /usr/local/bin/app /usr/local/bin/app

CMD ["app"]