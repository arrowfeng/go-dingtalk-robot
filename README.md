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
$ ./dingtalk --token xxx  send -f text
```
