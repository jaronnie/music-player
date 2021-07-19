FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct \
    PATH="/dist/bin:${PATH}"

WORKDIR /build

COPY . .

RUN go mod tidy && go build -o bin/music-player main.go

WORKDIR /dist

RUN cp -r /build/bin .

RUN cp -r /build/conf .

EXPOSE 8083

ENTRYPOINT ["music-player"]

