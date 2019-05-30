#!/bin/bash

kubectl config set-cluster ws-${NAMESPACE} \
    --embed-certs=true \
    --server=https://35.205.229.158 \
    --certificate-authority=${CA_PATH}
kubectl config set-credentials ws-${NAMESPACE} --token=${TOKEN}
kubectl config set-context ws-${NAMESPACE} \
    --cluster=ws-${NAMESPACE} \
    --user=ws-${NAMESPACE} \
    --namespace=${NAMESPACE}
kubectl config use-context ws-${NAMESPACE}
