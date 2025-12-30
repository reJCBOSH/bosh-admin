# bosh-admin

bosh-admin 是一个基于Go语言开发的后台管理系统。

## 项目结构

```
.
├── config                  # 配置文件目录
├── core                    # 核心模块
│   ├── ctx                 # 上下文相关
│   ├── db                  # 数据库相关
│   ├── exception           # 异常处理
│   └── log                 # 日志模块
├── global                  # 全局变量
├── initializer             # 初始化模块
├── middleware              # 中间件
├── migrations              # 数据库迁移
├── model                   # 数据模型
├── module                  # 功能模块
├── router                  # 路由配置
├── util                    # 工具类
├── websocket               # WebSocket相关
├── main.go                 # 入口文件
└── config.yaml             # 配置文件
```

## Docker 部署

本项目支持使用 Docker 进行容器化部署。

### 使用 Docker 直接构建

1. 构建镜像：
```bash
docker build -t bosh-admin .
```

2. 运行容器：
```bash
docker run -d -p 8089:8089 --name bosh-admin bosh-admin
```

### 使用 Docker Compose 部署（推荐）

项目提供了 `docker-compose.yml` 文件，可以一键部署包含数据库在内的完整服务：

```bash
docker-compose up -d
```

该命令将会启动以下服务：
- MySQL 数据库服务
- bosh-admin 应用服务

> 注意：首次启动时，MySQL 需要一定时间初始化，请稍等片刻再访问应用。

### 环境变量配置

可以通过修改 `config.yaml` 文件或者设置环境变量来配置服务：

- `SERVER_PORT`: 服务端口，默认为 8089
- 数据库相关配置请参考 `config.yaml` 文件

## 本地开发
