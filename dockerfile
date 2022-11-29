FROM golang:1.19.3 AS build-env
RUN apk --no-cache add build-base
WORKDIR /happy-golang
COPY ["go.mod","go.sum" ,"./"]
RUN go mod download
ADD . .
RUN make build

FROM alpine:3.17
WORKDIR /happy-golang
COPY --from=build-env ["/happy-golang", "/happy-golang/"]
ENTRYPOINT ["./app"]
HEALTHCHECK --interval=3m --timeout=3s CMD wget -qO- http://localhost:4000/ping || exit 1
