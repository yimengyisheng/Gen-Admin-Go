# Gen-Admin-Go

> **🚀 由 AI 驱动开发**
>
> 本项目由大语言模型 (LLM) 启动、构建并持续开发. 它是一个强大的示例, 展示了从初始架构设计、功能实现到持续优化的、由 AI 驱动的软件工程全过程.
>
> 整个开发流程是**对话式**的, 允许快速迭代, 并能够根据需求"生成"新的功能、增删改查接口和架构改进.

这是一个使用 Go 语言构建的现代化后端管理 API 项目, 遵循了清晰的、可扩展的项目架构.

## ✨ 功能特性

- **RESTful API:** 使用 `GET` 和 `POST` 提供清晰的 API 接口.
- **模块化架构:** 清晰分离的业务逻辑、数据访问和 API 展示层.
- **配置管理:** 使用 Viper 和 godotenv 进行灵活的环境配置.
- **数据库:** 使用 GORM 与 MySQL 进行交互.
- **身份认证:** 通过 JWT (JSON Web Tokens) 保护 API 路由.
- **日志系统:** 集成了强大的 Zap 日志库, 并封装了易于使用的辅助函数.
- **严格的参数校验:** 统一的参数绑定和校验层, 拒绝未知参数, 并提供友好的中文错误提示.

## 🛠️ 技术栈

- **Web 框架:** [Gin](https://github.com/gin-gonic/gin)
- **数据库 ORM:** [GORM](https://gorm.io/)
- **数据库驱动:** [MySQL Driver for GORM](https://gorm.io/docs/connecting_to_the_database.html#MySQL)
- **配置管理:** [Viper](https://github.com/spf13/viper)
- **环境变量:** [godotenv](https://github.com/joho/godotenv)
- **身份认证:** [JWT for Go](https://github.com/golang-jwt/jwt)
- **日志:** [Zap](https://github.com/uber-go/zap)
- **参数校验:** [Validator](https://github.com/go-playground/validator)

## 📂 项目结构

```
/ai_admin_project/
├── cmd/server/             # 应用启动入口
├── config/                 # Viper 配置
├── internal/
│   ├── handler/            # HTTP Handler (Gin)
│   ├── middleware/         # 中间件 (JWT认证)
│   ├── model/              # GORM 数据模型
│   ├── repository/         # 数据访问层 (GORM)
│   ├── request/            # 请求参数校验
│   ├── response/           # 响应数据封装
│   └── service/            # 业务逻辑层
├── pkg/
│   ├── auth/               # JWT 生成和解析
│   ├── logger/             # Zap 日志封装
│   └── utils/              # 通用工具函数 (参数校验, 翻译器)
├── go.mod                  # Go 模块管理
├── go.sum
├── .env                    # 环境变量 (模板)
└── README.md               # 项目说明文档
```

## 🚀 快速开始

### 1. 先决条件

- [Go](https://golang.org/dl/) (版本 >= 1.18)
- [MySQL](https://www.mysql.com/)

### 2. 克隆项目

```bash
git clone https://github.com/your_username/ai_admin_project.git
cd ai_admin_project
```

### 3. 配置环境变量

项目通过根目录下的 `.env` 文件进行配置. 您可以复制 `.env.example` (如果提供) 或手动创建它.

```bash
# .env 文件内容

# 服务器配置
SERVER_PORT=8080

# 数据库配置
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password # 替换为您的数据库密码
DB_NAME=ai_admin

# JWT 配置
JWT_SECRET=your_super_secret_key # 强烈建议修改为一个随机的复杂字符串
JWT_EXPIRATION=72 # Token 有效期 (小时)
```

### 4. 数据库设置

请确保您的 MySQL 服务正在运行, 并创建一个名为 `ai_admin` 的数据库.

```sql
CREATE DATABASE ai_admin CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 5. 安装依赖并运行

```bash
# 下载所有依赖项
go mod tidy

# 启动服务
go run cmd/server/main.go
```

服务启动后, 您将在控制台看到类似以下的输出, 表示服务已在 `8080` 端口上运行:

```
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] Listening and serving HTTP on :8080
```

## 📖 API 端点

我们提供了一个 Postman 集合文件 `postman_collection.json`, 您可以将其导入 Postman 以方便地测试所有 API.

### 用户 (Users)

- **注册:** `POST /register`
- **登录:** `POST /login`
- **获取个人信息:** `GET /api/profile` (需要认证)

### 商品 (Products)

所有商品相关的 API 都需要认证.

- **创建商品:** `POST /api/products/create`
- **获取商品列表 (分页):** `POST /api/products/list`
- **获取单个商品:** `POST /api/products/get`
- **更新商品:** `POST /api/products/update`
- **删除商品:** `POST /api/products/delete`

## 🤝 贡献

我们欢迎所有形式的贡献! 如果您有任何想法、建议或发现了 bug, 请随时提交 [Issues](https://github.com/your_username/ai_admin_project/issues).

如果您想贡献代码, 请遵循以下步骤:

1.  Fork 本项目.
2.  创建您的特性分支 (`git checkout -b feature/AmazingFeature`).
3.  提交您的更改 (`git commit -m 'Add some AmazingFeature'`).
4.  将您的分支推送到远程 (`git push origin feature/AmazingFeature`).
5.  开启一个 Pull Request.

## 📄 许可证

本项目使用 [MIT 许可证](LICENSE) 开源.
