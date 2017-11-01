FROM golang:1.9 as builder
WORKDIR /go/src/github.com/smarthut/automata
COPY . .
RUN make vendor
RUN make build

FROM alpine:3.6
COPY --from=builder /go/src/github.com/smarthut/automata/automata /
EXPOSE 8080
VOLUME ["/data"]
ENTRYPOINT ["/automata"]
