FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/chauhanr/shipcon/user-cli

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -o user-cli -a -installsuffix cgo .


FROM debian:latest

#RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/chauhanr/shipcon/user-cli/user-cli .

ENTRYPOINT ["./user-cli"]

CMD ["./user-cli"]