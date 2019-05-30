#!/bin/bash

kubectl config set-cluster ws-${K8S_USER} \
    --embed-certs=true \
    --server=https://35.205.229.158 \
    --certificate-authority=${CA_PATH}
kubectl config set-credentials ws-${K8S_USER} --token=${TOKEN}
kubectl config set-context ws-${K8S_USER} \
    --cluster=ws-${K8S_USER} \
    --user=ws-${K8S_USER} \
    --namespace=${K8S_USER}
kubectl config use-context ws-${K8S_USER}
