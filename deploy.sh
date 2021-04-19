#!/bin/bash

set -e

DEPLOYNAME="merchant-qr"
REPOSITORY="asia.gcr.io/krungthai-cmmn-dev/tungngern/merchant-qr-service"
IMAGE="$REPOSITORY:uat"
HELM="helm"
NAMESPACE="tungngern"

echo "Are you sure to deploy below services \n$DEPLOYNAME\nto k8s namespace: tungngern "
read -p "Please confirm [Y/N]? " -n 1 -r
echo # (optional) move to a new line
if [[ $REPLY =~ ^[Yy]$ ]]; then

  read -p "Do you want to build code in this repo before deploy [Y/N]? " -n 1 -r
  echo # (optional) move to a new line
  if [[ $REPLY =~ ^[Yy]$ ]]; then
    docker build -t $IMAGE .
    docker push $IMAGE
  fi

  echo "Connecting to k8s cluster cmmn uat"
  gcloud container clusters get-credentials ktb-cmmn-uat-gke-microservice --region asia-southeast1 --project krungthai-cmmn-uat
  echo "=============================================="
  helm upgrade --install --set name=$DEPLOYNAME-halfhalf $DEPLOYNAME-halfhalf -n $NAMESPACE $HELM \
    --set image.tag=uat \
    --set image.repository=$REPOSITORY \
    -f $HELM/values-qr-halfhalf.yaml
  echo "=============================================="
  helm upgrade --install --set name=$DEPLOYNAME-rcn $DEPLOYNAME-rcn -n $NAMESPACE $HELM \
    --set image.tag=uat \
    --set image.repository=$REPOSITORY \
    -f $HELM/values-qr-rcn.yaml
  echo "=============================================="
  helm upgrade --install --set name=$DEPLOYNAME-rptg $DEPLOYNAME-rptg -n $NAMESPACE $HELM \
    --set image.tag=uat \
    --set image.repository=$REPOSITORY \
    -f $HELM/values-qr-rptg.yaml
  echo "=============================================="
  helm upgrade --install --set name=$DEPLOYNAME-promptpay $DEPLOYNAME-promptpay -n $NAMESPACE $HELM \
    --set image.tag=uat \
    --set image.repository=$REPOSITORY \
    -f $HELM/values-qr-promptpay.yaml

fi
