## Gitlab 钉钉机器人通知

背景：钉钉上自带的Gitlab机器人太难用了，显示的信息少，且不会自动艾特指定人。

<img width="481" alt="image" src="https://github.com/AGou-ops/gitlab-dingtalk-webhook/assets/57939102/5fe408b1-abbf-4a17-aa1d-7106557bd828">

<img width="470" alt="image" src="https://github.com/AGou-ops/gitlab-dingtalk-webhook/assets/57939102/99dad75a-32a1-463d-aeae-4a89130aa2f4">

<img width="465" alt="image" src="https://github.com/AGou-ops/gitlab-dingtalk-webhook/assets/57939102/25040b37-46ed-41fe-aabb-b09cdb7b7538">


### 快速开始

预先准备：

钉钉机器人配置：

<img width="735" alt="image" src="https://github.com/AGou-ops/gitlab-dingtalk-webhook/assets/57939102/bcdd361c-7d27-48c6-8756-bdda3efece6a">

gitlab配置（记得勾上这些`trigger`，还有填好`Secret token`，token在`main.go`18行）：

<img width="990" alt="image" src="https://github.com/AGou-ops/gitlab-dingtalk-webhook/assets/57939102/ed5d9cb8-5569-4785-88ca-423df7bc8a6d">


使用Docker镜像快速开始：

```bash
docker run -ti --rm -p 8787:8787 -e WB_PATH=/webhooks \
  -e TOKEN=eb1axxxxxxx \
  -e SECRET=SECxxxxxxxxxxx \
  suofeiya/gitlab-dingtalk-webhook:latest
```

本地快速运行：

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

或者自己构建镜像：

```bash
docker build -t gitlab-dingtalk:v1.0 .
# 使用配置文件
docker run -d --restart always \
  --name gitlab_dingtalk_webhook \
  -v `pwd`/.env:/.env \
  -p 8787:8787 \ gitlab-dingtalk:v1.0
  
# 使用环境变量
docker run -d --restart always \
  --name gitlab_dingtalk_webhook \
  -e WB_PATH=/webhooks -e TOKEN=xxxx -e SECRET=SECxxxxx \
  -p 8787:8787 \
  gitlab-dingtalk:v1.0
```
