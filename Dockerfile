# Stage 1. Build.

FROM golang:1.22 as build

WORKDIR /app

COPY . ./

RUN go mod download

RUN make build

# Stage 2.

FROM debian:12-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates

RUN update-ca-certificates

COPY --from=build /app/apiserver /apiserver

EXPOSE 8888

CMD ["/apiserver"]
