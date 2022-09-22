FROM golang:1.18-alpine
WORKDIR /day-7
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
EXPOSE 8080
