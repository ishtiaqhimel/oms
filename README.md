# OMS (Order Management System)

Welcome to the OMS (Order Management System) application. In the following sections you will get idea about the whole system in brief:

## Services

- **Auth Server :** Users can signup or login. Users need to signup at first. Then they will get access token each new login.
- **Order Management System Server :** Simple order CRUD application.
- **Proxy :** Used `kong` as a public facing piece. All the traffic will be landed here. Auth termination is done here with the help of [JWT Plugin](https://docs.konghq.com/hub/kong-inc/jwt/) and also route check are done here.

## Install

You need a kubernetes cluster having loadbalancer implemneted. If you are using kind then follow [this](https://kind.sigs.k8s.io/docs/user/ingress/). You can deploy the whole application on Kubernetes cluster using Helm package manager.

At first, you need to clone the repository. Deploy the application using Helm by following command:

```bash
$ cd ./charts
$ helm install oms ./oms --namespace oms --create-namespace
```

To learn more about `oms` chart, see [here](./charts/oms/README.md).

See [demo](#demo) section to try out.

## Demo

You can try out the apis. The apis documentations can be found in `swagger.yaml` file on respective server's directory ([auth-server](./auth-server/swagger.yaml) or [oms-server](./oms-server/swagger.yaml)). You can use [swagger editor](https://editor.swagger.io/) to read documentations on UI. If you make any changes, use `make swagger` command to update swagger file.

At first, you need to signup:
```bash
$ curl -X POST "http://<address>/auth/signup" \
       -d '{"username": "Alice", "password": "12345"}'
```

To generate authentication token:
```bash
$ curl -X POST "http://<address>/auth/login" \
       -d '{"username": "Alice", "password": "12345"}' \
       -v
```

Here, you will get the `Bearer Token` in the headers. Copy and save the token.

Now, call the order apis. Let's try create an order:

```bash
$ curl -X GET "http://<address>/api/order" \
       -H "Authorization: <bearer-token>" \
       -d '{"username":"Alice", "productname":"tea"}'
```

Let's try list the orders:

```bash
$ curl -X GET "http://<address>/api/order" \
       -H "Authorization: <bearer-token>"
```

Here, you will see the list of orders.