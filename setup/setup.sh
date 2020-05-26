#!/bin/bash



# echo a message to standard error (used for messages not intended
# to be parsed by scripts, such as usage messages, warnings or errors)
echo_stderr() {

  echo "$@" >&2
  
}


if [[ ! $(which docker) ]]; then
	echo_stderr "Docker is not installed!"
	exit 3
fi


if [[ ! $(which minikube) ]]; then
	echo_stderr "Minikube is not installed!"
	exit 3

fi


if [[ ! $(which kubectl) ]]; then
	echo_stderr "Kubectl is not installed!"
	exit 3

fi


if $[[ $(minikube start) ]];
	echo "Environment is setup!"
	exit 1
fi

