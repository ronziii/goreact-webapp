SHELL := /bin/bash 

setup: 
	cd setup && ./setup.sh

status:
	cd deploy && ./deploy.sh --status

build-frontend:
	cd deploy && ./deploy.sh --build-frontend

build-backend:
	cd deploy && ./deploy.sh --build-backend

build-database:
	cd deploy && ./deploy.sh --build-database

build-all: 
	cd deploy && ./deploy.sh --build-all

deploy-frontend:
	cd deploy && ./deploy.sh --deploy-frontend

deploy-backend:
	cd deploy && ./deploy.sh --deploy-backend

deploy-database:
	cd deploy && ./deploy.sh --deploy-database

deploy-all:
	cd deploy && ./deploy.sh --deploy-all

backup-db:
	cd backup-db && ./backup_db.sh

help:
	cd deploy && ./deploy.sh --help



