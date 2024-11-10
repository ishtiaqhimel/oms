## OMS server 

CRUD apis for order management.

### TL;DR;

**Installing Chart**

```bash
$ helm upgrade -i oms-server ./oms-chart --namespace oms --create-namespace
```

**Uninstalling Chart**

```bash
$ helm uninstall oms-server -n oms
```

### Dependency Charts

- [kong](https://github.com/Kong/charts)

### Configuration

The following table lists the configurable parameters of the `oms-server` chart and their default values.

| Parameter | Description | Default |
|---|---|---|
| nameOverride | Overrides name template | `""` |
| fullnameOverride | Overrides fullname template | `""` |
| replicaCount | Number of oms-server replicas to create | `1` |
| image.repository | Name of oms-server container image | `ishtiaq99/oms-server` |
| image.pullPolicy | Container image pull policy | `Always` |
| image.tag | OMS server container image tag | `latest` |
| podAnnotations | Annotations passed to server pod(s). | `{}` |
| podLabels | Labels passed to server pod(s). | `component: oms` |
| service.type | Service type for oms server | `ClusterIP` |
| service.port | Service port for oms server | `3000` |
| ingress.annotations | Annotations passed to ingress | `konghq.com/strip-path: "true"` |
| ingress.className | | `kong` |
| ingress.host | | `""` | 
| env.dbName | Env variable passed to server pod(s). Represents the database name | `postgres` |
| env.dbUser | Env variable passed to server pod(s). Represents the database user name | `postgres` |
| env.dbPort | Env variable passed to server pod(s). Represents the database port | `"5432"` |
| env.dbHost | Env variable passed to server pod(s). Represents the database host name | `postgres` |
| secret.dbPassword | `DB_PASSWORD` key set in secret | `"12345"` |
| secret.jwtAlgo | `algorithm` key set in secret | `HS256` |
| secret.jwtSecret | `secret` key set in secret | `randomsecret` |
| secret.jwtKey | `key` key set in secret | `randomkey` |
| kong.enabled | | `true` |
| kong.ingressController.installCRDs | | `false` |

- Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`.
- Alternatively, a YAML file that specifies the values for the parameters can be provided using `--value` flag while installing the chart.