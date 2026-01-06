# Docker 环境变量配置说明

## 概述
本项目支持通过环境变量覆盖 config.yaml 中的数据库配置，以便在 Docker 环境中灵活管理配置。

## 环境变量说明

### 数据库配置
- `DB_HOST`: 数据库主机地址 (默认: 从 config.yaml 读取)
- `DB_PORT`: 数据库端口 (默认: 从 config.yaml 读取)
- `DB_USER`: 数据库用户名 (默认: 从 config.yaml 读取)
- `DB_PASSWORD`: 数据库密码 (默认: 从 config.yaml 读取)
- `DB_NAME`: 数据库名称 (默认: 从 config.yaml 读取)

### MySQL 容器配置
- `MYSQL_ROOT_PASSWORD`: MySQL root 用户密码
- `MYSQL_USER`: MySQL 普通用户 (不能是root)
- `MYSQL_PASSWORD`: MySQL 普通用户密码
- `MYSQL_DATABASE`: 初始化数据库名

## 使用方法

### 1. 本地开发
直接运行应用，将使用 config.yaml 中的配置：

```bash
go run main.go
```

### 2. Docker Compose 部署
使用 .env 文件管理所有配置：

```bash
# 修改 .env 文件以更新配置
# 然后启动服务
docker-compose up -d
```

MySQL 容器将使用环境变量自动创建用户和数据库，无需额外的初始化脚本。

### 3. 环境变量优先级
环境变量的优先级高于 config.yaml 文件中的配置。当环境变量存在时，应用将使用环境变量的值。

## 配置同步
当需要更新数据库密码等敏感信息时：
1. 只需修改 .env 文件中的相应变量
2. docker-compose.yml 会自动使用新的值
3. 应用容器会通过环境变量获取到新的数据库配置

> 注意：MYSQL_USER 不能设置为 "root"，因为 root 是特殊用户，只能通过 MYSQL_ROOT_PASSWORD 设置。

这样就实现了在单个位置修改配置的目标。