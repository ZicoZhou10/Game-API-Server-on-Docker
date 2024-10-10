# Game-API-Server-on-Docker
## 项目简介

这个项目是一个基于微服务架构的游戏 API 服务器，使用 Docker 进行容器化部署。它包含了玩家管理、房间管理、挑战系统、日志收集和支付处理等功能，旨在为 OXO 系列游戏提供后端支持。

## 技术栈

- 后端: Go (Gin 框架)
- 数据库: PostgreSQL
- 容器化: Docker
- API 网关: 自定义 Go 服务
- 消息队列: RabbitMQ (如果使用的话)
- 缓存: Redis (如果使用的话)

## 项目结构
APIServer/
├── api_gateway/
├── player_management/
├── room_management/
├── challenge_system/
├── log_collector/
├── payment_system/
└── docker-compose.yml

每个子目录都包含一个独立的微服务，具有自己的 API、模型和数据库连接。

## 安装和运行

1. 确保您的系统已安装 Docker 和 Docker Compose。

2. 克隆此仓库：
   git clone https://github.com/ZicoZhou10/Game-API-Server-on-Docker.git
   cd Game-API-Server-on-Docker

3. 运行服务：
   docker-compose up -d

4. 服务将在以下端口启动：
- API 网关: http://localhost:8080
- 玩家管理: http://localhost:8081
- 房间管理: http://localhost:8082
- 挑战系统: http://localhost:8083
- 日志收集: http://localhost:8084
- 支付处理: http://localhost:8085

## API 文档

### 玩家管理

- `POST /players`: 创建新玩家
- `GET /players`: 获取所有玩家
- `GET /players/{id}`: 获取特定玩家
- `PUT /players/{id}`: 更新玩家信息
- `DELETE /players/{id}`: 删除玩家

### 房间管理

- `POST /rooms`: 创建新房间
- `GET /rooms`: 获取所有房间
- `GET /rooms/{id}`: 获取特定房间
- `PUT /rooms/{id}`: 更新房间信息
- `DELETE /rooms/{id}`: 删除房间

### 挑战系统

- `POST /challenges`: 创建新挑战
- `GET /challenges/results`: 获取挑战结果

### 日志收集

- `POST /logs`: 添加新日志
- `GET /logs`: 获取日志（支持查询参数）

### 支付处理

- `POST /payments`: 处理新支付
- `GET /payments/{id}`: 获取支付信息

## 测试

您可以使用 curl 或 Postman 来测试 API。以下是一些示例命令：

### 玩家管理

1. 创建新玩家：
   curl -X POST http://localhost:8080/players -H "Content-Type: application/json" -d '{"name":"Test Player","level":1}'

2. 获取所有玩家：
   curl -X GET http://localhost:8080/players

3. 获取特定玩家（假设 ID 为 1）：
   curl -X GET http://localhost:8080/players/1

4. 更新玩家信息：
   curl -X PUT http://localhost:8080/players/1 -H "Content-Type: application/json" -d '{"name":"Updated Player","level":2}'

5. 删除玩家：
   curl -X DELETE http://localhost:8080/players/1

### 房间管理

1. 创建新房间：
   curl -X POST http://localhost:8080/rooms -H "Content-Type: application/json" -d '{"name":"Game Room 1","description":"A fun game room","status":"available"}'

2. 获取所有房间：
   curl -X GET http://localhost:8080/rooms

3. 获取特定房间（假设 ID 为 1）：
   curl -X GET http://localhost:8080/rooms/1

### 挑战系统

1. 创建新挑战：
   curl -X POST http://localhost:8080/challenges -H "Content-Type: application/json" -d '{"player_id":1,"amount":20.01}'

2. 获取挑战结果：
   curl -X GET http://localhost:8080/challenges/results

### 日志收集

1. 添加新日志：
   curl -X POST http://localhost:8080/logs -H "Content-Type: application/json" -d '{"player_id":1,"action":"login","details":"Player logged in"}'

2. 获取日志（带查询参数）：
   curl -X GET "http://localhost:8080/logs?player_id=1&action=login"

### 支付处理

1. 处理新支付：
   curl -X POST http://localhost:8080/payments -H "Content-Type: application/json" -d '{"player_id":1,"amount":50.00,"payment_method":"credit_card","payment_details":"4111111111111111"}'
   
2. 获取支付信息（假设支付 ID 为 1）：
   curl -X GET http://localhost:8080/payments/1

注意：在实际测试中，请确保替换示例中的 ID（如玩家 ID、房间 ID 等）为实际存在的 ID。

这些示例提供了对每个主要功能的基本 CRUD 操作的测试。

如果您的 API 有任何特殊的认证要求（如 JWT 令牌），您可能需要在命令中添加相应的头部。例如：

curl -X GET http://localhost:8080/players -H "Authorization: Bearer YOUR_JWT_TOKEN"


MIT License
