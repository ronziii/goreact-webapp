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
	build_frontend
    ;;

    --build-backend)
    build_backend
    ;;

    --build-all)
	build_all
	;;

	--deploy-frontend)
	deploy_frontend
	;;

    --deploy-backend)
	deploy_backend
	;;

	--deploy-database)
	deploy_database
	;;

	--deploy-all)
	deploy_all
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
    echo_stderr "Unknown option $1"
    show_usage
    exit 1
    ;;
esac
done




