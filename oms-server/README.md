## OMS Server

Documentation for OMS Server.

---

API Endpoints

| Endpoint | Function | Method | StatusCode |
| -------- | -------- | ------ | ---------- |
| `/order` | ListOrder | GET | Success - 200, Failure - 500 |
| `/order` | CreateOrder | POST | Success - 201, Failure - 400 |
| `/order/{id}` | DeleteOrder | DELETE | Success - 204, Failure - 404 |
| `/order/{id}` | GetOrderByID | GET | Success - 200, Failure - 404 |
| `/order/{id}` | UpdateOrderByID | PUT | Success - 200, Failure - 400, 404 |

[**NOTE :** need to add `/api` prefix for each endpoints if installed via provided [Helm chart](../charts/oms/)]

---

Data Model

```go
type Order struct {
	ID uint `gorm:"primaryKey"`
	Username string `gorm:"not null"`
	ProductName string `gorm:"not null"`
}
```

---

Build Docker Image

```bash
$ export REGISTRY=<your-registry-name>
$ export BUILD_PLATFORM=<platform-arch> # default: linux/amd64,linux/arm64
$ make build
```