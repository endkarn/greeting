FROM golang:1.12 as builder
WORKDIR /module
COPY ./go.mod /module/go.mod
COPY ./go.sum /module/go.sum
RUN go mod download
COPY . /module
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/app ./cmd/main.go
ENV GIN_MODE release
EXPOSE 3000
CMD ./app