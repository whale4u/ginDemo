FROM golang:alpine
LABEL MAINTAINER=whale4u
ENV TZ=Asia/Shanghai
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /app
COPY . .
RUN go build .

EXPOSE 8090
ENTRYPOINT ["./ginDemo"]