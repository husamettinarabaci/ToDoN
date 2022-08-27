
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

# ToDoN

Web Based ToDo List Application

##

![App Screenshot](./Documents/Images/webpage.png)

## Demo

Demo uygulama buradan erişebilirsiniz :

- [@live-demo](http://aab6a856cd8984348b4e34a58610665f-1150680336.eu-central-1.elb.amazonaws.com)

## Environment Variables

### Actions

To run this project, you will need to add the following environment variables to your 

Gtihub.com -> Repo -> Settings -> Secrets->Actions

`AWS_ACCESS_KEY_ID`

`AWS_REGION`

`AWS_SECRET_ACCESS_KEY`

`DOCKER_HUB_ACCESS_TOKEN` -> Docker.com -> Security > New Access Token

`DOCKER_HUB_USERNAME`

`KUBE_CONFIG_DATA` -> cat $HOME/.kube/config | base64

### Web Server - Test

To test this project, you will need to add the following environment variables

`MEMCACHE_SERVER_IP`

`MEMCACHE_SERVER_PORT` 

## Installation - K8S

### Namespaces

```bash
- kubectl create -f Configs/K8S/Namespaces/todon-prod-namespace.yaml
- kubectl create -f Configs/K8S/Namespaces/todon-test-namespace.yaml
```

### Config Maps

```bash
- kubectl create -f Configs/K8S/Configs/prod/memcacheserver-prod-config.yaml
- kubectl create -f Configs/K8S/Configs/prod/webserver-prod-config.yaml 
- kubectl create -f Configs/K8S/Configs/test/memcacheserver-test-config.yaml 
- kubectl create -f Configs/K8S/Configs/test/webserver-test-config.yaml 
```

### Cluster IPs

```bash
- kubectl create -f Configs/K8S/ClusterIPs/prod/memcacheserver-prod-clusterip.yaml 
- kubectl create -f Configs/K8S/ClusterIPs/test/memcacheserver-test-clusterip.yaml 
```

### LoadBalancers

```bash
- kubectl create -f Configs/K8S/LoadBalancers/prod/memcacheserver-prod-loadbalancer.yaml 
- kubectl create -f Configs/K8S/LoadBalancers/prod/webserver-prod-loadbalancer.yaml 
- kubectl create -f Configs/K8S/LoadBalancers/test/memcacheserver-test-loadbalancer.yaml 
- kubectl create -f Configs/K8S/LoadBalancers/test/webserver-test-loadbalancer.yaml 
```

### Deployments

```bash
- kubectl create -f Configs/K8S/Deployments/prod/memcacheserver-prod-deployment.yaml
- kubectl create -f Configs/K8S/Deployments/prod/webserver-prod-deployment.yaml 
- kubectl create -f Configs/K8S/Deployments/test/memcacheserver-test-deployment.yaml 
- kubectl create -f Configs/K8S/Deployments/test/webserver-test-deployment.yaml 
```

## Run Locally

Clone the project

```bash
  git clone https://github.com/husamettinarabaci/ToDoN.git
```

Go to the Mem-Cache Server project directory

```bash
  cd ToDoN/Services/BE_Services/memcacheserver/
```

Install dependencies

```bash
  go get
```

Build & Start the Mem-Cache Server

```bash
  go build
  ./memcacheserver
```

Go to the Web Server project directory

```bash
  cd ToDoN/Services/FE_Services/webserver/
```

Install dependencies

```bash
  go get
```

Build & Start the Web Server

```bash
  go build
  ./webserver
```

See [ToDoN](http://localhost)

## API Reference

#### Get all Todos

```http
  GET /api/v1/all
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| ` N/A   ` | `      ` |                            |

#### Add a new todo

```http
  POST /api/v1/add/
```

| Data      | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `Item`    | `string` | **Required**.                     |

## Running Tests

Go to the Mem-Cache Server project directory or Web Server project directory

```bash
  cd ToDoN/Services/BE_Services/memcacheserver/
  or
  cd ToDoN/Services/FE_Services/webserver/
```

To run tests, run the following command

```go
  go run test -v 
```

## Deployment

To deploy this project run

```bash
  git tag v*.*.* (etc. v2.0.0)
  git push --tags
```

After this action, the following events occur automatically:
- Docker Image Build && Push : (DOCKER_HUB_USERNAME)/memcacheserver_test:v2.0.0
- Docker Image Build && Push : (DOCKER_HUB_USERNAME)/memcacheserver_prod:v2.0.0
- Docker Image Build && Push : (DOCKER_HUB_USERNAME)/webserver_test:v2.0.0
- Docker Image Build && Push : (DOCKER_HUB_USERNAME)/webserver_prod:v2.0.0
- Aws - Eks Update : memcacheserver-prod-deployment
- Aws - Eks Update : webserver-prod-deployment 
- Aws - Eks Update : memcacheserver-test-deployment 
- Aws - Eks Update : webserver-test-deployment 


## CI/CD Pipeline

![CI/CD](./Documents/Images/pipeline.png)

![CI/CD](./Documents/Images/pipeline_done.png)

## Directory Structure

└── ToDoN

    ├── Configs

    │   ├── DockerFiles

    │   │   ├── Dockerfile_memcacheserver

    │   │   ├── Dockerfile_webserver

    │   │   └── Dockerfile_webserver_apicdctest

    │   ├── K8S

    │   │   ├── ClusterIPs
    │   │   │   ├── prod
    │   │   │   │   └── memcacheserver-prod-clusterip.yaml
    │   │   │   └── test
    │   │   │       └── memcacheserver-test-clusterip.yaml
    │   │   ├── Configs
    │   │   │   ├── prod
    │   │   │   │   ├── memcacheserver-prod-config.yaml
    │   │   │   │   └── webserver-prod-config.yaml
    │   │   │   └── test
    │   │   │       ├── memcacheserver-test-config.yaml
    │   │   │       └── webserver-test-config.yaml
    │   │   ├── Deployments
    │   │   │   ├── prod
    │   │   │   │   ├── memcacheserver-prod-deployment.yaml
    │   │   │   │   └── webserver-prod-deployment.yaml
    │   │   │   └── test
    │   │   │       ├── memcacheserver-test-deployment.yaml
    │   │   │       └── webserver-test-deployment.yaml
    │   │   ├── LoadBalancers
    │   │   │   ├── prod
    │   │   │   │   ├── memcacheserver-prod-loadbalancer.yaml
    │   │   │   │   └── webserver-prod-loadbalancer.yaml
    │   │   │   └── test
    │   │   │       ├── memcacheserver-test-loadbalancer.yaml
    │   │   │       └── webserver-test-loadbalancer.yaml
    │   │   └── Namespaces
    │   │       ├── todon-prod-namespace.yaml
    │   │       └── todon-test-namespace.yaml
    │   └── Tests
    │       └── webserver_apicdctest_config.apib
    ├── Documents
    │   ├── Images
    │   │   ├── configmaps.png
    │   │   ├── deployments.png
    │   │   ├── endpoints.png
    │   │   ├── memcacheserver_test_duplicate_data_result.png
    │   │   ├── memcacheserver_test_result.png
    │   │   ├── nodes.png
    │   │   ├── pipeline_done.png
    │   │   ├── pipeline.png
    │   │   ├── pods.png
    │   │   ├── services.png
    │   │   ├── webpage.png
    │   │   └── webserver_test_result.png
    │   ├── Tasks
    │   │   └── technologist.pdf
    │   └── Workflows_test
    │       ├── memcacheserver_audit.yml
    │       ├── memcacheserver_prod_build.yml
    │       ├── memcacheserver_prod_deploy.yml
    │       ├── memcacheserver_test_build.yml
    │       ├── memcacheserver_test_deploy.yml
    │       ├── memcacheserver.yml
    │       ├── webserver_audit.yml
    │       ├── webserver_prod_build.yml
    │       ├── webserver_prod_deploy.yml
    │       ├── webserver_test_build.yml
    │       ├── webserver_test_deploy.yml
    │       └── webserver.yml
    ├── go.work
    ├── go.work.sum
    ├── LICENSE
    ├── Note.md
    ├── README.md
    ├── README_screens.md
    └── Services
        ├── BE_Services
        │   └── memcacheserver
        │       ├── doc.go
        │       ├── go.mod
        │       ├── go.sum
        │       ├── memcacheserver.go
        │       └── memcacheserver_test.go
        ├── FE_Services
        │   └── webserver
        │       ├── doc.go
        │       ├── go.mod
        │       ├── go.sum
        │       ├── views
        │       │   └── index.html
        │       ├── webserver.go
        │       └── webserver_test.go
        └── Shareds
            └── proto
                └── item
                    ├── go.mod
                    ├── go.sum
                    ├── item_grpc.pb.go
                    ├── item.pb.go
                    └── item.proto

## Screenshots

[Screenshots](./README_screens.md)

## Tech Stack

![Go](https://img.shields.io/badge/Go-v1.19-blue)
![gRPC](https://img.shields.io/badge/gRPC-proto-blue)
![Docker](https://img.shields.io/badge/Docker-passing-green)
![Kubernetes](https://img.shields.io/badge/Kubernetes-MicroServices-blue)
![Aws](https://img.shields.io/badge/Aws-Eks-blue)
![github](https://img.shields.io/badge/Github-Actions-green)
![API](https://img.shields.io/badge/API-http-blue)
![Test](https://img.shields.io/badge/Test-unit-green)
![Test](https://img.shields.io/badge/Test-cdc-green)
![CI/CD](https://img.shields.io/badge/CI%20CD-automation-green)

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Contributing

Contributions are always welcome!

Please adhere to this project's `code of conduct`.

## Authors

- [@husamettinarabaci](https://www.github.com/husamettinarabaci)

