FROM golang:1.22.4-alpine AS bob-constructor

WORKDIR /

COPY . .
RUN go build -o ./bin/app ./cmd/main.go

FROM scratch AS runner

WORKDIR /

COPY --from=bob-constructor /bin/app app

CMD ["./app"]

