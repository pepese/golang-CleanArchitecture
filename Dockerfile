FROM golang:1.15.6-alpine3.13 as build
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ca-app

FROM alpine:3.13
RUN apk add --no-cache tzdata
ENV TZ Asia/Tokyo
COPY --from=build /build/ca-app /usr/local/bin/ca-app
EXPOSE 8080
CMD ["/usr/local/bin/ca-app", "server", "--type", "gin"]