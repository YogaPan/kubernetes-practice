FROM golang:1.23.1-bookworm AS builder
RUN mkdir /src
ADD . /src
WORKDIR /src

RUN go env -w GO111MODULE=auto
RUN go build -o main .

FROM gcr.io/distroless/base-debian12

WORKDIR /
COPY --from=builder /src/main /main
EXPOSE 8080
ENTRYPOINT ["/main"]