#!/bin/bash

set -e
export DOCKER_IMAGE_NAME=arkedemy-golang-course-app
docker image tag haidlir/${DOCKER_IMAGE_NAME} gcr.io/${PROJECT_NAME_PRD}/${DOCKER_IMAGE_NAME}:$TRAVIS_COMMIT

echo $GCLOUD_SERVICE_KEY_PRD | base64 --decode -i > ${HOME}/gcloud-service-key.json
gcloud auth activate-service-account --key-file ${HOME}/gcloud-service-key.json

gcloud --quiet config set project $PROJECT_NAME_PRD
gcloud --quiet config set container/cluster $CLUSTER_NAME_PRD
gcloud --quiet config set compute/zone ${CLOUDSDK_COMPUTE_ZONE}
gcloud --quiet container clusters get-credentials $CLUSTER_NAME_PRD

gcloud auth configure-docker
docker push gcr.io/${PROJECT_NAME_PRD}/${DOCKER_IMAGE_NAME}

yes | gcloud container images add-tag gcr.io/${PROJECT_NAME_PRD}/${DOCKER_IMAGE_NAME}:$TRAVIS_COMMIT gcr.io/${PROJECT_NAME_PRD}/${DOCKER_IMAGE_NAME}:latest

kubectl config view
kubectl config current-context

kubectl set image deployment.apps/${KUBE_DEPLOYMENT_NAME} ${KUBE_DEPLOYMENT_CONTAINER_NAME}=gcr.io/${PROJECT_NAME_PRD}/${DOCKER_IMAGE_NAME}:$TRAVIS_COMMIT -n production