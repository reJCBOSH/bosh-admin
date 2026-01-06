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

项目支持通过环境变量覆盖 `config.yaml` 中的数据库配置，实现配置的统一管理。

1. 配置环境变量：修改 `.env` 文件中的数据库账号密码等配置
2. 启动服务：`docker-compose up -d`

配置文件中的数据库连接信息会通过环境变量动态覆盖，只需在 `.env` 文件中修改一次即可。

可用的环境变量包括：
- `DB_HOST`: 数据库主机地址
- `DB_PORT`: 数据库端口  
- `DB_USER`: 数据库用户名
- `DB_PASSWORD`: 数据库密码
- `DB_NAME`: 数据库名称
- `MYSQL_ROOT_PASSWORD`: MySQL root 用户密码
- `MYSQL_USER`: MySQL 普通用户（不能是 "root"）
- `MYSQL_PASSWORD`: MySQL 普通用户密码
- `MYSQL_DATABASE`: 初始化数据库名

> 注意：`MYSQL_USER` 不能设置为 "root"，因为 root 是特殊用户，只能通过 `MYSQL_ROOT_PASSWORD` 设置。

## 本地开发
