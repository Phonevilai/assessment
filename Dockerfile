FROM golang:1.19-alpine as build-base

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/myapi .

#======================

FROM alpine:3.17

COPY --from=build-base /app/out/myapi /app/myapi
ENV TZ="Asia/Bangkok"

CMD ["/app/myapi"]