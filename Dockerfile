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
RUN go build -o bin/wechat-server
RUN sed -E "s/^run_mode.*/run_mode = \"release\"/g" conf/httpserver.toml

FROM scratch as prod

WORKDIR /app

COPY --from=0 /build/bin/ bin/
COPY --from=0 /build/conf/ conf/

EXPOSE 8080

CMD ["./bin/wechat-server"]
