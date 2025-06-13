# datastore-zero

A simple key-value store API built in Go with persistent storage and Kubernetes-native deployment using PVCs.

This project is designed as a minimal, production-adaptable example of:
- Writing a RESTful API in Go
- Using SQLite (or Postgres) as a backend
- Persisting data via PVCs in Kubernetes
- Building, containerizing, and deploying a service with full CI/CD support

---

## 🔧 Features

- ✅ Simple key/value CRUD API
- ✅ SQLite backend (Postgres-ready)
- ✅ Dockerized and K8s-ready
- ✅ Metrics endpoint for Prometheus
- 🛡️ Optional: Auth headers, rate limiting, and tracing

---

## 📦 API Endpoints

| Method | Path           | Description            |
|--------|----------------|------------------------|
| PUT    | `/store/:key`  | Store or update value  |
| GET    | `/store/:key`  | Retrieve value by key  |
| DELETE | `/store/:key`  | Delete value by key    |

_Optional future:_
- `GET /store/` – list all keys
- `POST /store/` – bulk insert from JSON

---

## 🗃️ Data Model

```go
type Record struct {
    Key   string `json:"key" gorm:"primaryKey"`
    Value string `json:"value"`
}


---

🐳 Running Locally (dev)

\\\ bash
go run cmd/main.go
\\\

Default config:
    •   Port: 8080
    •   DB: ./data/data.db

---
🚀 Deployment Plan
1.  Dockerize the app
2.  Mount /data to a PVC in Kubernetes
3.  Configure environment with ConfigMaps/Secrets
4.  Push metrics to Prometheus
5.  Add CI/CD GitHub Actions for build/lint/test/deploy
