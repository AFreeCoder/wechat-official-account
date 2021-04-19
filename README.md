## 项目用途

微信公众号号后端服务

## 项目结构

```
├── README.md
├── build.sh
├── conf
│   ├── httpserver.toml
│   └── wechat.toml
├── data
├── document
├── global
│   └── global.go
├── go.mod
├── go.sum
├── httpserver
│   ├── controller
│   │   ├── base.go
│   │   ├── get.go
│   │   ├── message
│   │   │   ├── handle.go
│   │   │   ├── message.go
│   │   │   └── template.go
│   │   └── post.go
│   ├── router.go
│   └── server.go
├── library
│   └── util
│       └── verify.go
├── logger
│   ├── logfield.go
│   └── logger.go
├── main.go
└── model
    └── qabot
        └── qabot.go
```