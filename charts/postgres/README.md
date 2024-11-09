## Postgres

Database

### TL;DR;

**Installing Chart**

```bash
$ helm upgrade -i postgres ./postgres --namespace postgres --create-namespace
```

**Uninstalling Chart**

```bash
$ helm uninstall postgres -n postgres
```

### Configuration

The following table lists the configurable parameters of the `postgres` chart and their default values.

| Parameter | Description | Default |
|---|---|---|
| nameOverride | Overrides name template | `""` |
| fullnameOverride | Overrides fullname template | `""` |
| replicaCount | Number of postgres replicas to create | `1` |
| image.repository | Name of postgres container image | `postgres` |
| image.pullPolicy | Container image pull policy | `IfNotPresent` |
| image.tag | Postgres container image tag | `latest` |
| podAnnotations | Annotations passed to server pod(s). | `{}` |
| podLabels | Labels passed to server pod(s). | `component: db` |
| resources | Resource limit and request | `requests: storage:1G` |
| service.port | Service port for server | `5432` |
| secret.db | `POSTGRES_DB` key set in secret | `postgres` |
| secret.user | `POSTGRES_USER` key set in secret | `postgres` |
| secret.dbPassword | `POSTGRES_PASSWORD` key set in secret | `"12345"` |

- Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`.
- Alternatively, a YAML file that specifies the values for the parameters can be provided using `--value` flag while installing the chart.