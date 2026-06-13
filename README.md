# gin-mini-agent

极简化的 Gin 脚手架，专为 Agent 开发场景设计。

## 特性

- **极简设计**：只保留核心功能，无冗余代码
- **最佳实践**：遵循 Gin 框架推荐的代码组织方式
- **开箱即用**：配置管理、日志系统、定时任务、参数校验已集成
- **优雅关闭**：支持平滑停机

## 项目结构

```
gin-mini-agent/
├── main.go                    # 应用入口
├── initialize/                # 初始化模块
│   ├── config.go             # 配置加载
│   ├── logger.go             # 日志初始化
│   ├── router.go             # 路由注册
│   ├── cron.go               # 定时任务初始化
│   └── validate.go           # 参数校验器
├── api/                       # API 处理器
│   └── health.go             # 健康检查
├── models/                    # 数据模型
│   └── resp.go               # 统一响应结构
├── pkg/                       # 公共包
│   ├── global/               # 全局变量
│   └── utils/                # 工具函数
├── cronjob/                   # 定时任务
├── conf/                      # 配置文件
├── logs/                      # 日志目录
├── go.mod
└── Dockerfile
```

## 快速开始

### 环境要求

- Go 1.25+
- Redis (可选)

### 安装运行

```bash
# 克隆项目
git clone https://github.com/your-repo/gin-mini-agent.git
cd gin-mini-agent

# 安装依赖
go mod tidy

# 运行
go run main.go
```

### 健康检查

```bash
curl http://localhost:8080/health
```

## 核心依赖

| 依赖 | 用途 |
|------|------|
| [gin](https://github.com/gin-gonic/gin) | Web 框架 |
| [viper](https://github.com/spf13/viper) | 配置管理 |
| [lumberjack](https://github.com/natefinch/lumberjack) | 日志轮转 |
| [cron](https://github.com/robfig/cron) | 定时任务 |
| [validator](https://github.com/go-playground/validator) | 参数校验 |

## 配置说明

配置文件位于 `conf/` 目录：

- `config.se.yml` - 开发环境
- `config.st.yml` - 测试环境
- `config.prd.yml` - 生产环境

通过环境变量 `RunMode` 切换环境：

```bash
export RunMode=prd
go run main.go
```

## API 端点

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/health` | 健康检查 |

## 扩展指南

### 添加新 API

1. 在 `api/` 目录创建处理器文件
2. 在 `initialize/router.go` 注册路由

```go
// api/user.go
package api

func GetUser(c *gin.Context) {
    c.JSON(200, gin.H{"id": c.Param("id")})
}

// initialize/router.go
r.GET("/user/:id", api.GetUser)
```

### 添加定时任务

1. 在 `cronjob/` 目录创建任务文件
2. 在 `initialize/cron.go` 注册任务

```go
// cronjob/my_task.go
type MyTask struct{}

func (t *MyTask) Run() {
    // 任务逻辑
}

// initialize/cron.go
c.AddJob("0 * * * * *", &cronjob.MyTask{})
```

## License

MIT
