## 项目用途

微信公众号号后端服务

## 项目结构

```
├── README.md
├── conf
│   └── httpserver.toml
├── data
├── document
├── go.mod
├── go.sum
├── httpserver
│   ├── controller
│   │   ├── base.go
│   │   ├── get.go
│   │   ├── message
│   │   │   ├── handle.go
│   │   │   └── message.go
│   │   └── post.go
│   ├── router.go
│   └── server.go
├── library
│   └── util
│       └── base.go
├── main.go
└── model
    └── text
        └── analysis.go
```