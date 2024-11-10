## Auth Server

Documentation for Auth Server.

---

API Endpoints

| Endpoint | Function | Method | StatusCode |
| -------- | -------- | ------ | ---------- |
| `/signup` | SignUp | POST | Success - 200, Failure - 400, 500 |
| `/login` | Login | POST | Success - 200, Failure - 400, 401, 500 |

[**NOTE :** need to add `/auth` prefix for each endpoints if installed via provided [Helm chart](../charts/oms/)]

---

Data Model

```go
type User struct {
	ID uint `gorm:"primaryKey" json:"-"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
```

---

Build Docker Image

```bash
$ export REGISTRY=<your-registry-name>
$ export BUILD_PLATFORM=<platform-arch> # default: linux/amd64,linux/arm64
$ make build
```