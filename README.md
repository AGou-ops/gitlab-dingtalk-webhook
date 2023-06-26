## Gitlab 钉钉机器人通知

背景：钉钉上自带的Gitlab机器人太难用了，显示的信息少，且不会自动艾特指定人。

### 快速开始

```bash
git clone https://github.com/AGou-ops/gitlab-dingtalk-webhook.git
cd gitlab-dingtalk-webhook
go mod tidy
# 在.env文件中将配置修改为你自己机器人的token和secret
cp .env.sample .env
go run .
```

切换监听端口：

```bash
go run . -p 9898
```

服务默认监听地址为: https://<YOUR_SERVER_IPADDR>:8787/webhooks

仅允许对URI为`/webhooks`的地址进行`POST`.

### 使用Docker运行

```bash
docker build -t gitlab-dingtalk:v1.0 .

docker run -d --restart always --name gitlab_dingtalk_webhook -p 8787:8787 gitlab-dingtalk:v1.0
```