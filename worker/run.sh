#!/bin/bash

readonly IMAGE=raonismaneoto/custom-lb-worker
readonly CONTAINER_NAME=custom-load-balancer-worker
readonly PROJECT_PATH="/go/src/github.com/raonismaneoto/CustomLB"

sudo docker pull $IMAGE
sudo docker stop $CONTAINER_NAME
sudo docker rm $CONTAINER_NAME
sudo docker run --name $CONTAINER_NAME -tdi $IMAGE

sudo docker exec $CONTAINER_NAME /bin/bash -c "$PROJECT_PATH/main" &