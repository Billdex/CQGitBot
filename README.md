#CQGitBot

---

##简介

CQGitBot是一个使用Golang开发，调用[cqhttp](https://github.com/richardchien/coolq-http-api)提供的酷Q的api接口开发的简单Github消息推送服务机器人。它在接收到Github的WebHook消息之后，会进行一些简单的处理并将消息转发到指定QQ群。

##使用说明

使用该机器人需要先下载安装[酷Q](https://cqp.cc/t/23253)和[CoolQ Http API](https://github.com/richardchien/coolq-http-api/releases)插件。安装好酷Q后，将下载好的cpk插件放入酷Q安装目录内的app目录下，即

```
|--CoolQAir
|  |--app
|     |--io.github.richardchien.coolqhttpapi.cpk
```

接下来启动酷Q并在应用管理中启用HTTP API插件即可开始使用。

Linux下可以使用cqhttp的Docker镜像[richardchien/cqhttp](https://hub.docker.com/r/richardchien/cqhttp/)

```
$docker pull richardchien/cqhttp:latest
$docker run -tid --rm --name cqhttp-cqgitbot \
 -v $(pwd)/coolq/cqgitbot:/home/user/coolq \
 -p 9000:9000 \ #noVNC端口，用于从浏览器管理酷Q
 -p 5700:5700 \ #HTTP API插件的端口
 -e COOLQ_ACCOUNT=你要登录的QQ号 \
 -e CQHTTP_POST_URL=http://127.0.0.1:7920/qq \
 -e CQHTTP_SERVE_DATA_FILES=yes \
 richardchien/cqhttp:latest
```

之后访问 `http://<你的IP>:9000/` 进入noVNC管理页面，默认密码为 `MAX8char` ,登录酷Q，即可开始使用。

最后，启动本项目的机器人，并在你想要推送消息的Github仓库中设置WebHook。

```
访问 http://仓库地址/settings/hooks
点击Add WebHook按钮添加
PayLoad URL 设为 http://<机器人所在IP>:7920/git
Content type 选择 application/json 
Secret 选填，作为鉴权用的密钥，可以为空
需要推送的Events事件可以选择 jsut push、全选、自选，根据需要进行选择 
```

添加好WebHook之后会发送一条ping消息，正常接收到该消息则说明配置正确且机器人正常启动。


