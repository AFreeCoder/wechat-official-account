FROM golang:alpine as builder

WORKDIR /build

COPY conf/ conf/
COPY global/ global/
COPY httpserver/ httpserver/
COPY library/ library/
COPY logger/ logger/
COPY model/ model/
COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go env -w GO111MODULE="on"
RUN go env -w GOPROXY="https://goproxy.cn,direct"
RUN mkdir bin
RUN sed -E "s/^run_mode.*/run_mode = \"release\"/g" conf/httpserver.toml
RUN go build -o bin/wechat-server

FROM alpine as prod

WORKDIR /app

COPY --from=builder /build/bin/ bin/
COPY --from=builder /build/conf/ conf/

EXPOSE 8080

CMD ["./bin/wechat-server"]
