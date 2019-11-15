go-dingtalk-robot
=================
    
这是一个基于 Go 语言实现的非常简单、易用、可扩展的钉钉机器人通知插件


## 前提

### 安装 Go 环境
请确保您已经正确安装了 Go 开发环境，建议安装 `Go version 1.12+`， 安装详情请参考[install instructions for Go](http://golang.org/doc/install.html)

#### 配置 GOPATH
请确保您的 `PATH` 路径下已经添加了`$GOPATH/bin`
```
$ export PATH=$PATH:$GOPATH/bin
```

#### 编译
在工程的当前目录下，进行编译

```
$ go build .
```

### 配置钉钉机器人
请确保您已经了解了钉钉机器人的相关信息以及相关配置，详情请参考[钉钉开发文档](https://ding-doc.dingtalk.com/doc#/serverapi2/krgddi)

#### 机器人配置
* 在钉钉群组里添加自定义机器人
* 配置安全选项（本项目支持自定义关键词、IP 地址段配置）
* 获取 access_token
* 为机器人设置头像、姓名（哈哈哈哈哈😂）

---

## 快速开始

当前自定义钉钉机器人支持文本 (text)、链接 (link)、markdown(markdown)、ActionCard、FeedCard 消息类型。

### Text 类型
```
$  ./dingtalk --token xxx text -c "我是一个没有感情的机器人" -a 188xxx8888 
```

### Link 类型
```
$  ./dingtalk --token xxx link -t "测试" -c "我是一个没有感情的机器人"  -p "https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png" -m "https://www.dingtalk.com"
```

### Markdown 类型
```
$ ./dingtalk --token xxx md -t "测试" -c "### 我是一个没有感情的机器人"
```

### ActionCard 类型
```
$ ./dingtalk --token xxx ac -t "测试" --stitle "阅读全文" --surl "https://www.dingtalk.com/" -c "### 我是一个没有感情的机器人"
```

### FeedCard 类型
```
$ ./dingtalk --token xxx fc -t "测试" -m "https://www.dingtalk.com/"  -p "https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png" 
```
### File 类型
可以发送任意类型的消息，前提是文件里的内容格式必须与钉钉开发文档的格式相匹配，例如 ./text.txt :
```txt
{
     "msgtype": "markdown",
     "markdown": {
         "title":"测试",
         "text": "## 我是一个没有感情的机器人 \n> 但是你喜欢我吗？\n\n> ![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png)\n> ###### [详情](http://www.dingtalk.com/) \n"
     },
    "at": {
        "atMobiles": [
        ],
        "isAtAll": false
    }
 }
```

```
$ ./dingtalk --token xxx send -f text
```

### 详细介绍
命令接收的参数可以通过两种形式注入，Options 和 Env。Options 优先级高于 Env

#### Text 命令
```
$ ./dingtalk --token xxx text [options]
```
##### Params 
|option|env|descripe|required|default|
|-|-|-|-|-|
|-c, --content|DINGTALK_CONTENT|设置发送内容|false|""|
|-f, --file|DINGTALK_FILE|从文件中读取发送内容|false|""|
|-a, --at|DINGTALK_ATMOBILES|@ 钉钉群组的某人, option 形式如 `-a 188xxxx9999` `-a 199xxxx9999`，env 形式如 `DINGTALK_ATMOBILES=188xxxx9999,199xxxx9999`|false|""|
|--all|DINGTALK_ISATALL|是否 @ 所有人|false|false|
|-h|查看帮助||false|""|
>**Tips -c, -f 二者不能都为空**

#### Link 命令
```
$ ./dingtalk --token xxx link [options]
```
##### Params 
|options|env|descripe|required|default|
|-|-|-|-|-|
|-t, --title|DINGTALK_TITLE|设置标题|true|""|
|-c, --content|DINGTALK_CONTENT|设置发送内容|false|""|
|-f, --file|DINGTALK_FILE|从文件中读取发送内容|false|""|
|-p, --purl|DINGTALK_PICTUREURL|设置图片 url|true|""|
|-m, --murl|DINGTALK_MESSAGEURL|设置消息 url|true|""|
>**Tips -c, -f 二者不能都为空**

#### Markdown 命令
```
$ ./dingtalk --token xxx md [options]
```
##### Params 
|options|env|descripe|required|default|
|-|-|-|-|-|
|-t, --title|DINGTALK_TITLE|设置标题|true|""|
|-c, --content|DINGTALK_CONTENT|设置发送内容, 可为 markdown 格式|false|""|
|-f, --file|DINGTALK_FILE|从文件中读取发送内容|false|""|
|-a, --at|DINGTALK_ATMOBILES|@ 钉钉群组的某人, option 形式如 `-a 188xxxx9999` `-a 199xxxx9999`，env 形式如 `DINGTALK_ATMOBILES=188xxxx9999,199xxxx9999`|false|""|
|--all|DINGTALK_ISATALL|是否 @ 所有人|false|false|

>**Tips -c, -f 二者不能都为空**

#### ActionCard 命令
```
$ ./dingtalk --token xxx ac [options]
```
##### Params 
|options|env|descripe|required|default|
|-|-|-|-|-|
|-t, --title|DINGTALK_TITLE|设置标题|true|""|
|-c, --content|DINGTALK_CONTENT|设置发送内容, 可为 markdown 格式|false|""|
|-f, --file|DINGTALK_FILE|从文件中读取发送内容|false|""|
|-a, --avatar|DINGTALK_ISHIDEAVATER|是否隐藏发送者头像, false 为不隐藏, true 为隐藏|false|false|
|-o, --orientation|DINGTALK_BTORIENTATION|设置按钮排列方式, false 为按钮竖直排列, true 为按钮横向排列|false|false|
|-m|DINGTALK_ISINDEPENDENT|是否为独立跳转|false|false|
|--stitle|DINGTALK_BTTITLE|按钮的名称。如果设置了 `-m`，可以设置多个，option 形式如 `--stitle` 感兴趣 `--stitle` 不感兴趣，env 形式如 `DINGTALK_BTTITLE=感兴趣,不感兴趣`。与 `--surl` 成对出现|true|""|
|--surl|DINGTALK_BTURL|设置点击按钮触发的 URL。如果设置了 `-m`，可以设置多个，option 形式如 `--surl "https://www.baidu.com"` `--surl "https://www.dingtalk.com"`，env 形式如 `DINGTALK_BTURL='"https://www.baidu.com","https://www.dingtalk.com"'`。与 `--stitle` 成对出现|true|""|

>**Tips -c, -f 二者不能都为空**

#### FeedCard 命令
```
$ ./dingtalk --token xxx fc [options]
```
##### Params 
|options|env|descripe|required|default|
|-|-|-|-|-|
|-t, --title|DINGTALK_TITLE|设置标题|true|""|
|-p, --purl|DINGTALK_PICTUREURL|设置图片 url|true|""|
|-m, --murl|DINGTALK_MESSAGEURL|设置消息 url|true|""|

>**Tips 上述三者必须成对出现**

---
### Todo
* 增加支持签名校验的安全选项
* 整合 Docker、Kubernetes
