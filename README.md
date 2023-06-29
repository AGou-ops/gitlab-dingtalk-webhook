## Gitlab 钉钉机器人通知

背景：钉钉上自带的Gitlab机器人太难用了，显示的信息少，且不会自动艾特指定人。

项目是练手的，使用了一些没必要用的框架，写的很一般。

<img width="436" alt="image" src="https://github.com/AGou-ops/gitlab-dingtalk-webhook/assets/57939102/7fa78be2-db55-4b60-be3e-a6f62a4ccc61">


<img width="492" alt="image" src="https://github.com/AGou-ops/gitlab-dingtalk-webhook/assets/57939102/1bea16a9-43be-421f-92b2-5cf9fe48dd32">



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

DockerHub: https://hub.docker.com/r/suofeiya/gitlab-dingtalk-webhook

```bash
docker build -t gitlab-dingtalk:v1.0 .
# 使用配置文件
docker run -d --restart always \
  --name gitlab_dingtalk_webhook \
  -v `pwd`/.env:/.env \
  -p 8787:8787 \
  gitlab-dingtalk:v1.0
  
# 使用环境变量
docker run -d --restart always \
  --name gitlab_dingtalk_webhook \
  -e WB_PATH=/webhooks -e TOKEN=xxxx -e SECRET=SECxxxxx \
  -p 8787:8787 \
  gitlab-dingtalk:v1.0
```
