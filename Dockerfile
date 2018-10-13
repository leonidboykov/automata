FROM golang:1.11 as builder
ENV GO111MODULE="on"
WORKDIR /go/src/github.com/smarthut/automata
COPY . .
RUN go build .

FROM alpine:latest
RUN apk --no-cache add tzdata zip ca-certificates
COPY --from=builder /go/src/github.com/smarthut/automata/automata /
EXPOSE 8080
VOLUME ["/scripts"]
ENTRYPOINT ["/automata"]
