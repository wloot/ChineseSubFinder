FROM node:17 as frontBuilder
WORKDIR /root/buildspace
COPY ./frontend .
RUN npm ci && npm run build

FROM golang:1.18-stretch as builder
ARG VERSION
WORKDIR /root/buildspace
COPY . .
COPY --from=frontBuilder /root/buildspace/dist/spa /root/buildspace/frontend/dist/spa
RUN go build -ldflags="-X main.AppVersion=${VERSION}" ./cmd/chinesesubfinder

FROM allanpk716/chinesesubfinder-base:latest
ENV TZ=Asia/Shanghai \
    PUID=1000 \
    PGID=1000 \
    DISPLAY=:99 \
    PS1="\u@\h:\w \$ "
COPY --from=builder /root/buildspace/chinesesubfinder /usr/local/bin/chinesesubfinder
COPY docker/full-rootfs /
ENTRYPOINT ["/init"]
