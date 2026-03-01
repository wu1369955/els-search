# Elasticsearch 搜索服务

## 项目结构

```
els-search/
├── backend/        # 后端代码
│   ├── api/        # API 层，处理 HTTP 请求和响应
│   ├── config/     # 配置管理
│   ├── repository/ # 数据访问层，处理 Elasticsearch 操作
│   ├── service/    # 业务逻辑层
│   ├── go.mod      # 依赖管理
│   ├── main.go     # 应用入口
│   └── README.md   # 项目说明
└── front/          # 前端代码
```

## 技术栈

- Go 1.20
- Gin 框架
- Elasticsearch 8.x

## 模块设计

### 1. 配置层 (config)
- 负责加载和管理应用配置
- 支持从环境变量读取配置

### 2. 数据访问层 (repository)
- 定义 Elasticsearch 操作接口
- 实现 Elasticsearch 搜索功能
- 与 Elasticsearch 客户端直接交互

### 3. 业务逻辑层 (service)
- 处理搜索业务逻辑
- 构建 Elasticsearch 查询
- 调用 repository 层执行搜索

### 4. API 层 (api)
- 处理 HTTP 请求和响应
- 解析查询参数
- 调用 service 层执行搜索

## 低耦合设计

- **依赖倒置**：使用接口定义各层之间的依赖关系
- **分层架构**：各层职责明确，边界清晰
- **依赖注入**：通过构造函数注入依赖，便于测试和替换

## 安装和运行

### 1. 安装依赖

```bash
go mod tidy
```

### 2. 配置 YAML 文件

编辑 `backend/config/config.yaml` 文件：

```yaml
server:
  addr: "localhost:8080"

elasticsearch:
  host: "http://localhost:9200"
```

### 3. 运行应用

```bash
cd backend
go run main.go
```

## API 接口

### 搜索接口

- **URL**: `/api/search`
- **方法**: GET
- **参数**:
  - `index`: Elasticsearch 索引名（必填）
  - `q`: 搜索关键词（必填）
  - `fields`: 搜索字段，多个字段用逗号分隔（可选）
  - `page`: 页码，默认 1（可选）
  - `size`: 每页大小，默认 10（可选）

- **示例请求**:
  ```
  GET /api/search?index=products&q=phone&fields=name,description&page=1&size=10
  ```

- **响应**:
  ```json
  {
    "hits": {
      "total": {
        "value": 10
      },
      "hits": [
        {
          "_source": {}
        }
      ]
    }
  }
  ```

### 健康检查接口

- **URL**: `/health`
- **方法**: GET
- **响应**:
  ```json
  {
    "status": "ok"
  }
  ```
