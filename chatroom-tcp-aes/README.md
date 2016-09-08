# 网络通信(服务端+客户端)

``` sh
├── bin		//跨平台可执行文件
│   ├── linux
│   │   ├── client
│   │   └── server
│   └── win64
│       ├── client.exe
│       └── server.exe
├── build.sh	//项目交叉编译脚本
├── client.go	//网络通信TCP客户端
├── server.go	//网络通信TCP服务端
└── README.md	//说明文件

3 directories, 8 files
```

# 加密算法——AES
``` go
// 此处设置16位密钥
var key = []byte("xxxxxxxxxxxxxxxx")
```

# 通信指令规则
- 指令
- 用户
- 消息
``` sh
# 设置昵称为ryanemax
nick ryanemax
# 发送消息“我们爱生活”给用户ryanemax
say ryanemax 我们热爱生活
```
