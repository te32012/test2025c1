FROM golang:latest AS build 

COPY . .

RUN go build -o /server ./cmd/main.go

FROM ubuntu:22.04 as production

COPY --from=build /server /app

CMD ["/app"]