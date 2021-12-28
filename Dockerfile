FROM golang:1.17.2 AS golang_builder
ENV GOPROXY=https://goproxy.cn,direct
ENV ROOT=/app
ENV CGO_ENABLED 0
RUN mkdir -p ${ROOT}
COPY ./server ${ROOT}/server
WORKDIR ${ROOT}/cmd/app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go_server

FROM node:14

COPY ./clinet ${ROOT}/clinet
WORKDIR ${ROOT}/clinet
ENV NODE_ENV=production
RUN npm install
RUN npm i -g serve
RUN npm run build
EXPOSE 5000
EXPOSE 22333
# 复制打包的Go文件到系统用户可执行程序目录下
COPY --from=golang_builder ${ROOT}/go_server ${ROOT}/go_server
RUN chmod +x ${ROOT}/go_server
# 容器启动时运行的命令
CMD ["serve", "-s", "dist"]
ENTRYPOINT ["${ROOT}/go_server"]
# FROM alpine:3.7
# # 配置国内源
# RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories
# RUN apk update
# RUN apk add ca-certificates
# # dns
# RUN echo "hosts: files dns" > /etc/nsswitch.conf
# RUN mkdir -p ${ROOT}

# WORKDIR ${ROOT}

