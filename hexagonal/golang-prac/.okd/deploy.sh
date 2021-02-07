#!/usr/bin/env sh

env > .env

echo "processing config map..."
oc process -f ConfigMap.yml --param-file=.env --ignore-unknown-parameters=true | oc apply -f -

echo "processing secrets..."
oc process -f Secret.yml --param-file=.env --ignore-unknown-parameters=true | oc apply -f -

echo "processing service..."
oc process -f Service.yml --param-file=.env --ignore-unknown-parameters=true | oc apply -f -

echo "processing routes..."
oc process -f Routes.yml --param-file=.env --ignore-unknown-parameters=true | oc apply -f -

echo "processing deployment config..."
oc process -f DeploymentConfig.yml --param-file=.env --ignore-unknown-parameters=true | oc apply -f -

echo "rolling out $SERVICE_NAME ..."
oc rollout latest dc/ganje

echo "rolling out $SERVICE_NAME-redis ..."
oc rollout latest dc/ganje-redis
