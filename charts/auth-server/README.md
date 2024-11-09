## Auth server 

- User SignUp and Login
- Generate JWT token during login

### TL;DR;

**Installing Chart**

```bash
$ helm upgrade -i auth-server ./auth-chart --namespace auth --create-namespace
```

**Uninstalling Chart**

```bash
$ helm uninstall auth-server -n auth
```

### Configuration

The following table lists the configurable parameters of the `auth-server` chart and their default values.

| Parameter | Description | Default |
|---|---|---|
| nameOverride | Overrides name template | `""` |
| fullnameOverride | Overrides fullname template | `""` |
| replicaCount | Number of auth-server replicas to create | `1` |
| image.repository | Name of auth-server container image | `ishtiaq99/auth-server` |
| image.pullPolicy | Container image pull policy | `Always` |
| image.tag | Auth server container image tag | `latest` |
| podAnnotations | Annotations passed to server pod(s). | `{}` |
| podLabels | Labels passed to server pod(s). | `component: auth` |
| service.type | Service type for auth server | `ClusterIP` |
| service.port | Service port for auth server | `3000` |
| ingress.annotations | Annotations passed to ingress | `konghq.com/strip-path: "true"` |
| ingress.className | | `kong` |
| ingress.host | | `""` | 
| env.dbName | Env variable passed to server pod(s). Represents the database name | `postgres` |
| env.dbUser | Env variable passed to server pod(s). Represents the database user name | `postgres` |
| env.dbPort | Env variable passed to server pod(s). Represents the database port | `"5432"` |
| env.dbHost | Env variable passed to server pod(s). Represents the database host name | `postgres` |
| secret.dbPassword | `DB_PASSWORD` key set in secret | `"12345"` |
| secret.jwtSecret | `JWT_SECRET` key set in secret | `randomsecret` |
| secret.jwtKey | `JWT_KEY` key set in secret | `randomkey` |

- Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`.
- Alternatively, a YAML file that specifies the values for the parameters can be provided using `--value` flag while installing the chart.