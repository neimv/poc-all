FROM golang:1.19 AS builder
LABEL environment="staging"

WORKDIR /go/src/github.com/neimv/poc-all/
COPY . .
RUN go get

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app .

FROM alpine:latest

ENV environment=$env

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/neimv/poc-all/app ./
# COPY --from=builder /go/src/github.com/neimv/poc-all/.env ./

# RUN export $(grep -v '^#' .env | xargs -d '\n')

EXPOSE 8000

CMD [ "./app" ]