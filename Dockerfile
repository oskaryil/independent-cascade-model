FROM golang:1.15 as builder

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -o icm .

FROM alpine:latest

RUN mkdir -p /usr/src/app

WORKDIR /usr/src/app

COPY --from=builder /go/src/app/icm /usr/src/app/
COPY --from=builder /go/src/app/android.csv /usr/src/app/

CMD ["./icm", "-f", "android.csv"]
