#!/bin/bash

build_frontend() {

	cd ../frontend || exit
	docker build -t goreact/frontend .
}


build_backend() {

	cd ../backend || exit
	docker build -t goreact/backend .
}

build_database() {

	cd ../database || exit
	docker build -t goreact/database .
}

build_all() {
	
	build_backend
	build_frontend	

}


deploy_backend() {

	cd ../kubernetes || exit
	kubectl create -f backend-deployment.yml
	kubectl create -f backend-service.yml
}

deploy_frontend() {

	cd ../kubernetes || exit
	kubectl create -f frontend-deployment.yml
	kubectl create -f frontend-service.yml
}


deploy_database() {

	cd ../database || exit
	kubectl create -f postgres-deployment.yml
	kubectl create -f postgres-service.yml

}

deploy_all() {

	deploy_database
	deploy_backend
	deploy_frontend
}


check_status() {
	kubectl get pods
}


show_usage() {

  echo_stderr "Usage: [make setup]|[make deploy-frontend]|[make deploy-backend]|[make deploy-all]|[make build-frontend]\n \
  			   [make build-backend]|[make build-database]|[make build-all]|[make status] "
}


# echo a message to standard error (used for messages not intended
# to be parsed by scripts, such as usage messages, warnings or errors)
echo_stderr() {

  echo "$@" >&2
  
}
