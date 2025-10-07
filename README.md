# Greatest Works - 分布式MMO游戏服务器

基于Go语言和领域驱动设计(DDD)架构开发的分布式大型多人在线游戏服务器，采用现代化微服务设计，支持高并发和分布式部署。

## 🎯 项目概述

这是一个企业级的分布式MMO游戏服务器项目，采用领域驱动设计(Domain-Driven Design)架构模式，提供高性能、可扩展、易维护的游戏服务器解决方案。项目采用分布式多节点架构，支持独立部署和扩展。

## ✨ 核心特性

- 🏗️ **DDD架构**: 采用领域驱动设计，清晰的架构分层和职责分离
- 🌐 **分布式设计**: 多节点独立部署，支持水平扩展
- 🚀 **高性能网络**: 基于Go原生RPC + TCP + HTTP多协议支持
- 🔧 **微服务架构**: 认证服务、网关服务、游戏服务独立部署
- 💾 **多数据库支持**: MongoDB + Redis 混合存储策略
- 🔐 **安全认证**: JWT认证系统，保障用户数据安全
- 🎮 **完整游戏功能**: 涵盖现代MMO游戏的核心系统
- 📊 **实时同步**: 高频率的游戏状态同步和事件处理
- 🛡️ **容错设计**: 完善的错误处理、监控和恢复机制
- 🐳 **容器化部署**: Docker和Kubernetes支持
- 📚 **完整文档**: 详细的API文档和架构说明

## 🏗️ 分布式架构设计

本项目采用分布式多节点架构，将游戏服务器拆分为三个独立的服务节点：

### 服务节点

#### 🔐 认证服务 (Auth Service)
- **协议**: HTTP
- **端口**: 8080
- **职责**: 用户认证、授权、会话管理
- **功能**: 登录、注册、令牌管理、权限控制

#### 🌐 网关服务 (Gateway Service)  
- **协议**: TCP
- **端口**: 9090
- **职责**: 客户端连接管理、协议转换、负载均衡
- **功能**: 连接管理、消息路由、协议转换

#### 🎮 游戏服务 (Game Service)
- **协议**: Go原生RPC
- **端口**: 8081
- **职责**: 核心游戏逻辑、领域模型、业务规则
- **功能**: 玩家管理、战斗系统、排行榜、社交系统

### 通信协议

```
客户端 ──HTTP──> 认证服务
  │
  └──TCP──> 网关服务 ──RPC──> 游戏服务
```

- **客户端 ↔ 认证服务**: HTTP (RESTful API)
- **客户端 ↔ 网关**: TCP (游戏协议)
- **网关 ↔ 游戏服务**: Go原生RPC (内部通信)
- **其他服务 ↔ 游戏服务**: Go原生RPC (服务间通信)

## 📁 项目结构

```
greatestworks/
├── cmd/                        # 应用程序入口
│   ├── auth-service/           # 认证服务
│   │   └── main.go
│   ├── gateway-service/        # 网关服务
│   │   └── main.go
│   └── game-service/           # 游戏服务
│       └── main.go
├── configs/                    # 配置文件
│   ├── auth-service.yaml       # 认证服务配置
│   ├── gateway-service.yaml    # 网关服务配置
│   ├── game-service.yaml       # 游戏服务配置
│   └── docker.yaml            # Docker环境配置
├── docs/                       # 项目文档
│   ├── api/                   # API文档
│   ├── architecture/          # 架构文档
│   └── diagrams/              # 架构图表
├── application/                # 应用层
│   ├── commands/              # 命令处理器
│   ├── handlers/              # 事件处理器
│   ├── queries/               # 查询处理器
│   └── services/              # 应用服务
├── internal/                   # 内部模块
│   ├── domain/                # 领域层
│   │   ├── player/           # 玩家领域
│   │   ├── battle/           # 战斗领域
│   │   ├── social/           # 社交领域
│   │   ├── building/         # 建筑领域
│   │   ├── pet/              # 宠物领域
│   │   ├── ranking/          # 排行榜领域
│   │   └── minigame/         # 小游戏领域
│   ├── infrastructure/        # 基础设施层
│   │   ├── persistence/      # 数据持久化
│   │   ├── cache/            # 缓存服务
│   │   ├── messaging/        # 消息服务
│   │   ├── network/          # 网络服务
│   │   ├── config/           # 配置管理
│   │   └── logging/          # 日志服务
│   └── interfaces/            # 接口层
│       ├── http/             # HTTP接口
│       ├── tcp/              # TCP接口
│       └── rpc/              # RPC接口
├── scripts/                    # 开发脚本
│   ├── start-services.bat     # Windows启动脚本
│   ├── start-services.sh      # Linux/Mac启动脚本
│   ├── build.sh              # 构建脚本
│   └── deploy.sh             # 部署脚本
├── docker-compose.yml          # Docker编排
├── Dockerfile                  # Docker镜像
├── Makefile                   # 构建工具
├── go.mod                     # Go模块定义
└── README.md                  # 项目说明
```

## 🛠️ 技术栈

### 核心技术
- **语言**: Go 1.21+
- **架构模式**: 领域驱动设计 (DDD) + 分布式架构
- **网络协议**: HTTP + TCP + Go原生RPC
- **数据库**: MongoDB (主数据库) + Redis (缓存)
- **消息队列**: NATS (可选)
- **认证**: JWT + 自定义认证
- **服务发现**: 支持Consul、Etcd等

### 开发工具
- **构建工具**: Make + Go Modules
- **容器化**: Docker + Docker Compose
- **编排**: Kubernetes
- **代码质量**: golangci-lint + 自定义规范
- **文档**: Markdown + 架构图

### 监控与运维
- **日志**: 结构化日志 + 分级输出
- **监控**: 自定义指标收集
- **健康检查**: HTTP健康检查接口
- **配置管理**: YAML配置 + 环境变量

## 🚀 快速开始

### 📋 环境要求

- **Go**: 1.21 或更高版本
- **MongoDB**: 4.4+ (推荐 5.0+)
- **Redis**: 6.0+ (推荐 7.0+)
- **Docker**: 20.10+ (可选，用于容器化部署)

### 📦 安装依赖

```bash
# 克隆项目
git clone https://github.com/phuhao00/greatestworks.git
cd greatestworks

# 安装Go依赖
go mod tidy
```

### ⚙️ 配置文件

项目使用独立的配置文件，每个服务都有自己的配置：

```bash
# 认证服务配置
cp configs/auth-service.yaml configs/auth-service-dev.yaml

# 网关服务配置  
cp configs/gateway-service.yaml configs/gateway-service-dev.yaml

# 游戏服务配置
cp configs/game-service.yaml configs/game-service-dev.yaml
```

### 🎮 启动服务

#### 方式一：使用启动脚本（推荐）

**Windows:**
```bash
scripts/start-services.bat
```

**Linux/Mac:**
```bash
./scripts/start-services.sh
```

#### 方式二：手动启动

```bash
# 启动认证服务
go run cmd/auth-service/main.go

# 启动游戏服务（新终端）
go run cmd/game-service/main.go

# 启动网关服务（新终端）
go run cmd/gateway-service/main.go
```

#### 方式三：Docker启动

```bash
# 启动完整环境
docker-compose up -d

# 查看服务状态
docker-compose ps
```

### 🔧 服务地址

启动后，各服务将在以下地址运行：

- **认证服务**: http://localhost:8080
- **游戏服务**: rpc://localhost:8081
- **网关服务**: tcp://localhost:9090

## 🏛️ DDD领域架构

### 核心领域 (Core Domains)

#### 🎮 玩家领域 (Player Domain)
- **职责**: 玩家基础信息、等级经验、属性管理
- **核心实体**: Player, PlayerStats, PlayerProfile
- **主要功能**: 玩家创建、升级、属性计算、状态管理

#### ⚔️ 战斗领域 (Battle Domain)
- **职责**: 战斗逻辑、技能系统、伤害计算
- **核心实体**: Battle, Skill, Damage, BattleResult
- **主要功能**: PvP/PvE战斗、技能释放、战斗结算

#### 🏠 社交领域 (Social Domain)
- **职责**: 聊天、好友、家族、队伍系统
- **核心实体**: Chat, Friend, Guild, Team, Mail
- **主要功能**: 社交互动、组队协作、消息通信

#### 🏗️ 建筑领域 (Building Domain)
- **职责**: 建筑系统、家园管理、建筑升级
- **核心实体**: Building, BuildingTemplate, BuildingUpgrade
- **主要功能**: 建筑建造、升级、功能管理

#### 🐾 宠物领域 (Pet Domain)
- **职责**: 宠物系统、宠物培养、宠物战斗
- **核心实体**: Pet, PetTemplate, PetSkill
- **主要功能**: 宠物获取、培养、进化、战斗辅助

#### 🏆 排行榜领域 (Ranking Domain)
- **职责**: 各类排行榜、积分统计、奖励发放
- **核心实体**: Ranking, RankingEntry, RankingReward
- **主要功能**: 排名计算、榜单更新、奖励分发

#### 🎯 小游戏领域 (Minigame Domain)
- **职责**: 各种小游戏、活动玩法、特殊奖励
- **核心实体**: Minigame, MinigameSession, MinigameReward
- **主要功能**: 小游戏逻辑、积分计算、奖励发放

## 🌐 网络协议设计

### 多协议支持
- **HTTP**: 认证服务，RESTful API
- **TCP**: 网关服务，游戏客户端连接
- **RPC**: 服务间通信，Go原生RPC

### TCP协议格式
```
+--------+--------+--------+----------+
| Magic  | Length | Type   | Data     |
| 2bytes | 4bytes | 2bytes | Variable |
+--------+--------+--------+----------+
```

### 消息分类
- **0x1xxx**: 系统消息 (登录、心跳、错误)
- **0x2xxx**: 玩家消息 (属性、状态、升级)
- **0x3xxx**: 社交消息 (聊天、好友、邮件)
- **0x4xxx**: 战斗消息 (技能、伤害、结果)
- **0x5xxx**: 建筑消息 (建造、升级、管理)
- **0x6xxx**: 宠物消息 (获取、培养、战斗)
- **0x7xxx**: 排行榜消息 (查询、更新、奖励)
- **0x8xxx**: 小游戏消息 (开始、操作、结算)
- **0x9xxx**: 管理消息 (GM命令、系统公告)

## 🗄️ 数据存储设计

### MongoDB 集合设计

#### 核心业务集合
- **players**: 玩家基础信息和状态
- **player_stats**: 玩家统计数据和属性
- **battles**: 战斗记录和结果
- **guilds**: 公会信息和成员关系
- **buildings**: 建筑数据和状态
- **pets**: 宠物信息和属性
- **rankings**: 排行榜数据和历史
- **minigames**: 小游戏记录和积分

### Redis 缓存策略

#### 热点数据缓存
- **在线玩家**: `online:players:{server_id}`
- **玩家会话**: `session:{player_id}`
- **排行榜**: `ranking:{type}:{period}`
- **公会信息**: `guild:{guild_id}`

#### 临时数据缓存
- **战斗状态**: `battle:{battle_id}`
- **队伍信息**: `team:{team_id}`
- **聊天频道**: `chat:{channel_id}`
- **活动状态**: `event:{event_id}`

## 👨‍💻 开发指南

### 🏗️ DDD开发模式

#### 添加新领域
1. 在 `internal/domain/` 下创建领域目录
2. 定义领域实体、值对象和聚合根
3. 实现领域服务和仓储接口
4. 创建对应的应用服务
5. 实现基础设施层的具体实现
6. 添加接口层的处理器

#### 领域开发规范
```go
// 领域实体示例
type Player struct {
    id       PlayerID
    name     string
    level    int
    exp      int64
    stats    PlayerStats
    // 领域行为
}

func (p *Player) LevelUp() error {
    // 领域逻辑实现
}
```

### 🔧 开发工具使用

#### Make命令
```bash
make setup      # 初始化开发环境
make dev        # 启动开发服务器
make build      # 构建生产版本
make test       # 运行测试
make lint       # 代码质量检查
make clean      # 清理构建产物
make docs       # 生成文档
```

### 📊 性能优化策略

#### 数据库优化
- **连接池管理**: 合理配置MongoDB和Redis连接池
- **索引优化**: 为查询频繁的字段创建合适索引
- **分片策略**: 大数据量集合采用分片存储
- **读写分离**: 读操作使用从库，写操作使用主库

#### 缓存策略
- **多级缓存**: 内存缓存 + Redis缓存 + 数据库
- **缓存预热**: 服务启动时预加载热点数据
- **缓存更新**: 采用Cache-Aside模式更新缓存
- **缓存穿透**: 使用布隆过滤器防止缓存穿透

#### 网络优化
- **连接复用**: TCP连接池和HTTP Keep-Alive
- **消息批处理**: 批量处理非实时消息
- **压缩传输**: 大数据包启用压缩
- **协议优化**: 使用二进制协议减少传输开销

## 🚀 部署指南

### 🐳 Docker部署

#### Docker Compose部署
```bash
# 启动完整环境（包含MongoDB、Redis）
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f
```

#### 单容器部署
```bash
# 构建镜像
docker build -t greatestworks .

# 运行认证服务
docker run -d --name auth-service -p 8080:8080 greatestworks auth-service

# 运行游戏服务
docker run -d --name game-service -p 8081:8081 greatestworks game-service

# 运行网关服务
docker run -d --name gateway-service -p 9090:9090 greatestworks gateway-service
```

### ☸️ Kubernetes部署

#### 基础部署
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: auth-service
  namespace: gaming
spec:
  replicas: 2
  selector:
    matchLabels:
      app: auth-service
  template:
    metadata:
      labels:
        app: auth-service
    spec:
      containers:
      - name: auth-service
        image: greatestworks:latest
        ports:
        - containerPort: 8080
        env:
        - name: SERVICE_TYPE
          value: "auth-service"
```

### 🔧 生产环境配置

#### 环境变量
```bash
# 服务配置
export SERVICE_TYPE="auth-service"  # auth-service, game-service, gateway-service
export SERVER_PORT=8080
export SERVER_HOST=0.0.0.0

# 数据库配置
export MONGODB_URI="mongodb://mongo-cluster:27017/gamedb"
export REDIS_ADDR="redis-cluster:6379"

# 认证配置
export JWT_SECRET="your-production-secret-key"

# 日志配置
export LOG_LEVEL=info
export LOG_FORMAT=json
```

## 📚 API文档

详细的API文档请参考：
- [REST API文档](docs/api/rest-api.md)
- [TCP协议文档](docs/api/tcp-protocol.md)
- [RPC接口文档](docs/api/rpc-api.md)

## 🏗️ 架构文档

深入了解系统架构：
- [DDD设计文档](docs/architecture/ddd-design.md)
- [分布式架构](docs/architecture/distributed-architecture.md)
- [数据库设计](docs/architecture/database-design.md)

## 🤝 贡献指南

我们欢迎所有形式的贡献！请阅读 [CONTRIBUTING.md](CONTRIBUTING.md) 了解详细信息。

### 贡献流程
1. **Fork** 项目到你的GitHub账户
2. **创建** 功能分支 (`git checkout -b feature/amazing-feature`)
3. **提交** 你的更改 (`git commit -m 'Add some amazing feature'`)
4. **推送** 到分支 (`git push origin feature/amazing-feature`)
5. **创建** Pull Request

### 开发规范
- 遵循 [Go代码规范](https://golang.org/doc/effective_go.html)
- 编写单元测试，保持测试覆盖率 > 80%
- 更新相关文档
- 通过所有CI检查

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 📞 联系我们

- **项目主页**: [https://github.com/phuhao00/greatestworks](https://github.com/phuhao00/greatestworks)
- **问题反馈**: [GitHub Issues](https://github.com/phuhao00/greatestworks/issues)
- **讨论交流**: [GitHub Discussions](https://github.com/phuhao00/greatestworks/discussions)

## 📈 项目状态

![Build Status](https://github.com/phuhao00/greatestworks/workflows/CI/badge.svg)
![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![Docker Pulls](https://img.shields.io/docker/pulls/greatestworks/server.svg)

## 🎯 路线图

### v2.0.0 (计划中)
- [ ] 服务网格集成
- [ ] GraphQL API支持
- [ ] 实时数据分析和BI
- [ ] 多语言客户端SDK
- [ ] 云原生部署优化

### v1.5.0 (开发中)
- [ ] 管理后台界面
- [ ] 性能监控面板
- [ ] 自动化测试覆盖
- [ ] 服务发现集成

### v1.0.0 ✅ (已发布)
- [x] 分布式架构重构完成
- [x] 多节点服务分离
- [x] Go原生RPC通信
- [x] Docker容器化支持
- [x] 基础监控和日志
- [x] 完整文档体系

---

<div align="center">

**⭐ 如果这个项目对你有帮助，请给我们一个Star！⭐**

*Built with ❤️ by the Greatest Works Team*

</div>