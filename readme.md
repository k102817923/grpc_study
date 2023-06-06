## grpc
> 微服务

**单体架构**
1. 一旦某个服务宕机, 会引起整个应用不可用, 隔离性差;
2. 只能整体应用进行伸缩, 浪费资源, 可伸缩性差;
3. 代码耦合在一起, 可维护性差

**微服务架构** 
解决了单体架构的弊端, 但同时引入了新的问题:
1. 代码冗余; 
2. 服务和服务之间存在调用关系; 
3. 服务拆分后, 服务和服务之间发生的是进程和进程之间的调用, 服务器和服务器之间的调用, 那么就需要发起网络调用;
4. 网络调用我们能立马想起的就是HTTP, 但是在微服务架构中, HTTP虽然便捷方便, 但性能较低;
5. 这时候就需要引入RPC(远程过程调用), 通过自定义协议发起TCP调用, 来加快传输效率

官网: `https://grpc.io/`
中文文档: `http://doc.oschina.net/grpc`

RPC的全称是Remote Procedure Call(远程过程调用), 这是一种协议, 是用来屏蔽分布式计算中的各种调用细节, 使得你可以像是本地调用一样直接调用一个远程的函数

**客户端与服务端沟通的过程**
1. 客户端发送数据(以字节流的方式);
2. 服务端接受并解析, 根据约定知道要执行什么, 然后把结果返回给客户

**RPC**
1. RPC就是将上述过程封装下, 使其操作更加优化;
2. 使用一些大家都认可的协议, 使其规范化;
3. 做成一些框架, 直接或间接产生利益

而GRPC又是什么呢?用官方的话来说:
`A high-performance open-source universal RPC framework`(GRPC是一个高性能的、开源的通用的RPC框架)

在RPC中, 我们称调用方为client, 被调用方为server。跟其他的RPC框架一样, GRPC也是基于"服务定义"的思想。
简单来说, 就是我们通过某种方式来描述一个服务, 这种描述方式是和语言无关的。
在这个"服务定义“的过程中, 我们描述了我们提供的服务的服务名是什么, 有哪些方法可以被调用, 这些方法有什么样的入参, 有什么样的回参。
也就是说, 在定义好了这些服务、这些方法之后, GRPC会屏蔽底层的细节, client只需要直接调用定义好的方法, 就能拿到预期的返回结果。
对于server端来说, 还需要实现我们定义的方法。
同样的, GRPC也会帮我们屏蔽底层的细节, 我们只需要实现所定义的方法的具体逻辑即可。
你可以发现, 在上面的描述过程中, 所谓的"服务定义", 和定义接口的语义是很接近的。
我更愿意理解为这是一种"约定", 双方约定好接口, 然后server实现这个接口, client调用这个接口的代理对象, 至于其他的细节, 则交给GRPC。
此外, GRPC还是语言无关的。你可以用C++作为服务端, 使用Golang、Java等作为客户端。
为了实现这一点, 我们在"定义服务"和编解码的过程中, 应该是做到语言无关的。

因此, GRPC使用了Protocol Buffers。
这是谷歌开源的一套成熟的数据结构序列化机制, 你可以把它当成一个代码生成工具以及序列化工具。
这个工具可以把我们定义的方法, 转换成特定语言的代码。
比如你定义了一种类型的参数, 它会帮你转换成Golang中的strut结构体, 你定义的方法, 它会帮你转换成func函数。
此外, 在发送请求和接受响应的时候, 这个工具还会完成对应的编解码工作, 将你即将发送的数据编码成GRPC能够传输的形式, 又或者将即将接收到的数据解码为编程语言能够理解的数据格式。

序列化: 将数据结构或对象转换成二进制串的过程
反序列化: 将在序列化过程中所产生的二进制串转换成数据结构或者对象的过程

protobuf是谷歌开源的一种数据格式, 适合高性能, 对响应速度有要求的数据传输场景。
因为profobur是二进制数据格式, 需要编码和解码, 数据本身不具有可读性, 因此只能反序列化之后得到真正可读的数据。
**优势**
1. 序列化后体积相比json和xml很小, 适合网络传输;
2. 支持跨平台多语言;
3. 消息格式升级和兼容性还不错;
4. 序列化反序列化速度很快

## 安装Protobuf
1. 下载protocol buffers: `https://github.com/protocolbuffers/protobuf/releases`
2. 配置环境变量: `D:\protoc-23.2-win64\bin`
3. 检查它是否有效, 打开cmd, 输入"protoc"命令
4. 安装GRPc的核心库: `go get google.golang.org/grpc`
5. 上面安装的是protocol编译器。它可以生成各种不同语言的代码。因此, 除了这个编译器之外, 我们还需要配合各个语言的代码生成工具。 
对于Golang来说, 称为protoc-gen-go。注意, `github.com/golang/protobuf/protoc-gen-go`和`google.golang.org/protobuf/cmd/protoc-gen-go`是不同的。
区别在于前者是旧版本, 后者是google接管后的新版本, 它们之间的API是不同的, 也就是说用于生成的命令以及生成的文件都是不一样的。 
因为目前的GRPC-go源码中的example用的是后者的生成方式, 所以我们也采取最新的方式。
6. 你需要安装两个库: 
   `go instal] google.golang.org/protobuf/cmd/protoc-gen-go@latest`
   `go insta1 google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
因为这些文件在安装GRPC的时候, 就已经下载下来了, 因此使用install命令就可以了, 而不需要使用get命令。然后查看`$GOWORKS/bin`路径下是否已经生成了两个新文件。

## proto文件编写
- 创建编写proto文件, client和server的内容是一样的
- 执行命令`protoc --go_out=. hello.proto` `protoc --go-grpc_out=. hello.proto`

## proto文件介绍
> message

- protobuf中定义一个消息类型式, 是通过关键字message字段指定的;
- 消息就是需要传输的数据格式的定义, message关键字类似于C++中的class、JAVA中的class、Go中的struct;
- 在消息中承载的数据分别对应于每一个字段, 其中每个字段都有一个名字和一种类型, 一个proto文件中可以定义多个消息类型

> 字段规则

- required: 消息体中必填字段, 不设置会导致编码异常。在protobuf2中使用, 在protobuf3中被删去;
- optional: 消息体中可选字段, protobuf3没有了required、optional等说明关键字, 都默认为optional;
- repeated: 消息体中可重复字段, 重复的值的顺序会被保留, 在Go中重复的会被定义为切片

> 消息号

在消息体的定义中, 每个字段都必须要有一个唯一的标识号, 标识号是[1, 2^29-1]范围内的一个整数。

> 嵌套消息

可以在其他消息类型中定义、使用消息类型, 在下面的例子中, person消息就是定义在PersonInfo消息内。
```protobuf
message PersonInfo {
  message Person {
    string name = 1;
    int32 height = 2;
    repeated int32 weight = 3;
  }
  repeated Person info = 1;
}
```
如果要在它的父消息类型的外部重用这个消息类型, 需要用PersonInfo.Person的形式使用它。
```protobuf
message PersonMessage {
  PersonInfo.Person info = 1;
}
```

> 服务定义

如果想要将消息类型用在RPC系统中, 可以在.proto文件中定义一个RPC服务接口, protocol buffer编译器将会根据所选择的语言生成服务接口代码及存根。
```protobuf
service SearchService {
  // rpc 服务函数名(参数) 返回(返回参数)
   rpc Search(SearchRequest) returns (SearchResponse)
}
```
上述定义了一个RPC服务, 该方法接收SearchRequest, 返回SearchResponse。

> 服务端编写
- 创建GRPC Server对象, 可以理解为是Server端的抽象对象
- 将Server(包含需要被调用的服务端接口)注册到GRPC Server的内部注册中心, 这样可以在接收到请求时, 通过内部的服务发现, 发现该服务端接口并进行逻辑处理
- 创建Listen, 监听TCP端口
- GRPC Server开始list.Accept, 直到Stop

> 客户端编写
- 创建与给定目标(服务端)的连接交互
- 创建Server的客户端对象
- 发送RPC请求, 等待同步响应, 得到回调后返回响应结果
- 输出响应结果

### 认证-安全传输
GRPC是一个典型的C/S模型, 需要开发客户端和服务端, 客户端和服务端需要达成协议, 使用某一个确认的传输协议来传输数据。GRPC通常默认是使用protobuf来作为传输协议, 也可以使用其他自定义的。
客户端在与服务端进行通信之前, 客户端需要知道自己的数据是发给哪一个服务端, 同理服务端也需要知道自己的数据要返回给谁, 这里则引申出GRPC的认证。
此处的认证, 指的不是用户身份的认证, 而是指多个server和多个client之间是如何识别对方并进行安全的数据传输。
- SSL/TLS的认证方式(采用HTTP2)
- 基于Token的认证方式(基于安全连接)
- 不采用任何措施的连接, 不安全的连接(默认采用HTTP1)
- 自定义的身份认证

### SSL/TLS认证方式
1. 通过openssl生成证书和私钥, 官网下载(`https://www.openssl.org/source`) or 安装包(`https://slproweb.com/products/Win32OpenSSL.html`)
2. 配置环境变量`OpenSSL-Win64\bin`
3. 命令行测试`openssl`

> 生成证书
```shell
# 1、生成私钥
openssl genrsa -out server.key 2048
# 2、生成证书, 全部回车即可, 可以不填
openssl req -new -x509 -key server.key -out server.crt -days 36500
# 3、生成csr
openssl req -new -key server.key -out server.csr

# 4、更改openssl.cnf(Linux下是openssl.cfg)
# 复制一份安装的openssl的bin目录里面的`openssl.cnf`文件到项目所在的目录
# 找到`[ CA_default ]`, 打开`copy_extensions = copy`(就是把前面的#去掉)
# 找到`[ req ]`, 打开`req_extensions = v3_reg`
# 找到`[ v3_req ]`, 添加`subjectAltName = @alt_names`
# 添加新的标签`[ alt_names ]`和标签字段`DNS.1 = *.grpcstudy.com`

# 5、生成证书私钥test.key
openssl genpkey -algorithm RSA -out test.key
# 通过私钥test.key生成证书请求文件test.csr(注意cfg和cnf)
openssl req -new -nodes -key test.key -out test.csr -days 3650 -subj "/C=cn/OU=myorg/O=mycomp/CN=myname" -config ./openssl.cfg -extensions v3_req
# test.csr是上面生成的证书请求文件, ca.crt/server.key是CA证书文件和key, 用来对test.csr进行签名认证。这两个文件在第一部分生成。
# 生成SAN证书pem
openssl x509 -req -days 365 -in test.csr -out test.pem -CA server.crt -CAkey server.key -CAcreateserial -extfile ./openssl.cfg -extensions v3_req
```

### Token认证
GRPC提供了接口PerRPCCredentials, 接口位于credentials包下, 接口含有两个方法GetRequestMetadata() & RequireTransportSecurity()。
第一个方法用于获取元数据信息, 即客户端提供的key, value键值对, context用于控制超时和取消, uri是请求入口处的uri。
第二个方法用于是否需要基于TLS认证进行安全传输, 如果返回值为true, 则必须加上TLS验证。
GRPC将各种认证方式浓缩统一到一个凭证(credentials)上, 可以单独使用一种凭证, 比如只使用TLS凭证或者只使用自定义凭证, 也可以多种凭证组合。
GRPC提供统一的API验证机制, 使研发人员使用方便, 这也是GRPC设计的巧妙之处。