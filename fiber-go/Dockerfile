FROM golang:latest as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o main .

FROM scratch

WORKDIR /app/

COPY --from=build /app/main ./main
COPY --from=build /app/.env.local /.env.local

CMD ["./main"]