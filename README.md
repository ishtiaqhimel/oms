# OMS (Order Management System)

Welcome to the OMS (Order Management System) application. In the following sections you will get idea about the whole system in brief:

## Services

- **Auth Server :** Users can signup or login. Users need to signup at first. Then they will get access token each new login.
- **Order Management System Server :** Simple order CRUD application.
- **Proxy :** Used `kong` as a public facing piece. All the traffic will be landed here. [Need to deploy on kubernetes cluster using Helm chart to test]

## Install

You can run the application by cloning this to your local machine or directly from Docker image. Also you can deploy the whole application on Kubernetes cluster using Helm package manager.

**Run on Local Machine**

Use `Postgres` as database.

```bash
$ docker run --name oms-postgres -e POSTGRES_PASSWORD=<password> -d postgres
```

To build your own image, use the following commands:

```bash
$ cd ./auth-server # use `cd ./oms-server` to build image for oms-server
$ export REGISTRY=<your-docker-registry>
$ export BUILD_PLATFORM=<your-preferred-arch> # ex. linux/amd64
$ make build
```

Now run the images providing necessary environment variables. You can find the variables list in `.env.example` file on respective server's directory ([auth-server](./auth-server/) or [oms-server](./oms-server/)).

Now you can try out the apis. The apis documentations can be found in `swagger.yaml` file on respective server's directory ([auth-server](./auth-server/swagger.yaml) or [oms-server](./oms-server/swagger.yaml)). You can use [swagger editor](https://editor.swagger.io/) to read documentations on UI. If you make any changes, use `make swagger` command to update swagger file.

**Run on Kubernetes Cluster**

Deploy the application using Helm by following command:

```bash
$ helm install oms ./oms --namespace oms --create-namespace
```

To learn more about `oms` chart, see [here](./charts/oms/README.md).

After successfully deploying, you will find an `ingress` object has been created. Here we have defined the rules for the http paths. Using the address in `ingress`, now you can try out the apis. 

## Demo

At first, you need to signup:
```bash
$ curl -X POST "http://<address>:<port>/auth/signup" \
       -d '{"username": "Alice", "password": "12345"}'
```

To generate authentication token:
```bash
$ curl -X POST "http://<address>:<port>/auth/login" \
       -d '{"username": "Alice", "password": "12345"}' \
       -v
```

Here, you will get the `Token` in the headers. Copy and save the token.

Now, call the order apis. Let's try list orders api:
```bash
curl -X GET "http://<address>:<port>/api/order" \
     -H "Token: <token>"
```

Here, you will see the list of orders.