GoReact-Webapp
=============

## Introduction

GoReact-Webapp is a simple application with a frontend, backend and a database designed to run in containers hosted on Minikube(Kubernetes for the local machine)

## Quickstart

### Prerequisites

* [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)
* [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl)
* [docker](https://docs.docker.com/engine/install/)

### Getting Started

As shown below, you may simply clone the GitHub repo to your local machine.

```shell
$ git clone https://github.com/ronnie420/goreact-webapp.git
```

The Kubernetes cluster configuration in this context will pull images from public repos rather than try to pull the images from local docker env/registry due to permissions.


## Usage

### Makefile

The commands/actions to build or deploy the application as shown below should be executed from the root folder where the Makefile is.

*** Check Environment Setup ***
```shell
$ make setup
```

*** Build All Applications ***
```shell
$ make build-all
```

*** Build Frontend Application ***
```shell
$ make build-frontend
```

*** Build Backend Application ***
```shell
$ make build-backend
```

*** Build Database (Docker Image) ***
```shell
$ make build-database
```
*** Deploy All Applications ***
```shell
$ make deploy-all
```

*** Deploy Frontend Application ***
```shell
$ make deploy-frontend
```

*** Deploy Backend Application ***
```shell
$ make deploy-backend
```

*** Deploy Backend Application ***
```shell
$ make deploy-backend
```

*** Show Help Menu ***
```shell
$ make help
	
```


*** Backup database script ***
This requires an IAM user with FullAccess S3

```shell
$ make backup-db
	
```
