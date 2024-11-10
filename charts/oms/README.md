## OMS

This chart deploys Order Management System on a Kubernetes cluster using Helm package manager.

### TL;DR;

**Installing Chart**

```bash
$ helm upgrade -i oms ./oms --namespace oms --create-namespace
```

**Uninstalling Chart**

```bash
$ helm uninstall oms -n oms
```

### Dependency Charts

- [auth-server](../auth-server/)
- [oms-server](../oms-server/)
- [postgres](../postgres/)

### Configuration

The following table lists the configurable parameters of the `oms` chart and their default values.

| Parameter | Description | Default |
|---|---|---|
| nameOverride | Overrides name template | `""` |
| fullnameOverride | Overrides fullname template | `""` | 
| postgres.enabled | | `true` |
| auth-server.enabled | | `true` |
| auth-server.service.port | | `3000` |
| auth-server.env.dbHost | | `oms-postgres` |
| oms-server.enabled | | `true` |
| oms-server.service.port | | `4000` |
| oms-server.env.dbHost | | `oms-postgres` |

[**NOTE:** The naming format for `auth-server.env.dbHost` & `oms-server.env.dbHost` is `<RELEASE_NAME>-postgres`.]

- Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`.
- Alternatively, a YAML file that specifies the values for the parameters can be provided using `--value` flag while installing the chart.