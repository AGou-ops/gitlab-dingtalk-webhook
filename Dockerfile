# build stage
# FROM harbor.nblh.local/library/golang:1.20.5-alpine AS build-stage
FROM golang:1.20.5-alpine AS build-stage

WORKDIR /app

COPY . .

# 设置个代理吧，不然拉都拉不下来...
ENV http_proxy=10.11.43.78:7890 https_proxy=10.11.43.78:7890

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -x -v -o /gitlab_dingtalk_amd64

# run server stage
# FROM harbor.nblh.local/library/alpine:latest
FROM alpine:latest

WORKDIR /

COPY .env .

COPY --from=build-stage /gitlab_dingtalk_amd64 /gitlab_dingtalk_amd64

EXPOSE 8787

ENTRYPOINT [ "/gitlab_dingtalk_amd64" ]
