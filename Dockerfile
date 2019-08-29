FROM golang:1.12 as builder
WORKDIR /module
COPY ./src/go.mod /module/go.mod
COPY ./src/go.sum /module/go.sum
RUN go mod download
COPY ./src/ /module
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/app ./cmd/greeting

FROM alpine
WORKDIR /root/
COPY --from=builder /module/bin .
ENV GIN_MODE release
EXPOSE 3000
CMD ./app