FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o stress-cli main.go
CMD [ "./stress-cli" ]

FROM scratch
COPY --from=builder /app/stress-cli .
ENTRYPOINT [ "./stress-cli" ]

