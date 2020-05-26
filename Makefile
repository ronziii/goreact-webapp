SHELL := /bin/bash 

status: 
	cd deploy && ./deploy.sh --status

build: 
	cd ../deploy && ./deploy.sh --status

