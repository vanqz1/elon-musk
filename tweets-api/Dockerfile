FROM golang:1.18-alpine

RUN apk add --no-cache git
WORKDIR /app
COPY go.mod .
COPY go.sum .
ENV GO111MODULE=on
RUN go mod download
COPY . .
RUN go build -o ./build/tweets .

EXPOSE 8080
CMD ["./build/tweets"]