# datastore-zero

A simple key-value store API built in Go with persistent storage and Kubernetes-native deployment using PVCs.

This project is designed as a minimal, production-adaptable example of:
- Writing a RESTful API in Go
- Using MariaDB as a backend
- Persisting data via PVCs in Kubernetes
- Building, containerizing, and deploying a service with full CI/CD support

---

## 🔧 Features

- ✅ Simple key/value CRUD API
- ✅ MariaDB backend
- ✅ Dockerized and K8s-ready
- ✅ Metrics endpoint for Prometheus
- 🧲 `/hire` route – because marketing is a feature too
- 🛡️ Optional: Auth headers, rate limiting, and tracing


---

## 📦 API Endpoints


| Method | Path           | Description            |
|--------|----------------|------------------------|
| PUT    | `/store/:id`   | Store or update value  |
| GET    | `/store/:id`   | Retrieve value by ID   |
| DELETE | `/store/:id`   | Delete value by ID     |
| GET    | `/store`       | List all records       |
| GET    | `/hire`        | Fun route that says "hire me!" |

---

---

## 🗃️ Data Model

```go
type Record struct {
    ID    string `json:"id" gorm:"primaryKey"`
    Value string `json:"value"`
}
```

---

## 🐳 Running Locally (dev)

```bash
go run cmd/main.go
```

Default config:
- Port: 8080
- DB: ./data/data.db

---

## 🚀 Deployment Plan
1. Dockerize the app
2. Mount /data to a PVC in Kubernetes
3. Configure environment with ConfigMaps/Secrets
4. Push metrics to Prometheus
5. Add CI/CD GitHub Actions for build/lint/test/deploy


--

## 🧲 Why "datastore-zero"?

Because every great engineer should ship something that's:
- Functional
- Minimal
- Tested
- Deployable

And ideally... gets them hired.

👋 [Let's talk](mailto:ericbotcher@gmail.com.com) — or just run `curl http://localhost:8080/hire`
