#!/bin/bash


source util-func.sh

if [[ ! $(which minikube) ]]; then
	echo "Minikube is not installed!"
fi


if [[ ! $(which kubectl) ]]; then
	echo "Kubectl is not installed!"
fi


if [[ ! $(which docker) ]]; then
	echo "Docker is not installed!"
fi



## Loop through the arguments passed
for i in "$@"
do
case $i in
    --build-frontend)
	echo "Build frontend!"
    ;;

    --build-backend)
    echo "Build backend!"
    ;;

    --build-all)
	echo "Build all!"
	;;

	--deploy-frontend)
	echo "Build frontend!"
	;;

    --deploy-backend)
	echo "Deploy backend!"
	;;

	--deploy-database)
	echo "Deploy database!"
	;;

	--deploy-all)
	echo "Deploy all!"
	;;

	--status)
	echo "Check status"
	check_status
	;;

	--help)
    show_usage
    exit 0
    ;;

    *)
    # unknown option
    echo "Unknown option $1"
    show_usage
    exit 1
    ;;
esac
done




