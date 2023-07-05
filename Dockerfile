FROM golang:1.19 AS builder

WORKDIR /go/src/github.com/alexellis/href-counter/
COPY go.mod .
COPY go.sum .
COPY main.go .
RUN go get

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/alexellis/href-counter/app ./

EXPOSE 8000

CMD ["./app"]