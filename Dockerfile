FROM golang:1.22-alpine3.16

ARG GO_LDFLAGS
RUN sed -e 's/dl-cdn[.]alpinelinux.org/mirrors.ustc.edu.cn/g' -i~ /etc/apk/repositories && apk add --update --no-cache gcc musl-dev
COPY ./ /build
WORKDIR /build
RUN CGO_ENABLED=1 GOPROXY="https://goproxy.cn,direct" go build -o . -ldflags="${GO_LDFLAGS} -w -s -extldflags -static" 

FROM alpine:latest
WORKDIR /app
COPY --from=0  /build/app /app/
# 安装tzdata软件包
RUN apk --no-cache add tzdata
# 设置时区为Asia/Shanghai（可以根据需要修改）
ENV TZ=Asia/Shanghai
EXPOSE 8000

CMD ["/app/app"]

