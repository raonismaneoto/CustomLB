#!/bin/bash

readonly IMAGE=raonismaneoto/custom-load-balancer
readonly CONTAINER_NAME=custom-load-balancer
readonly PROJECT_PATH="/go/src/github.com/raonismaneoto/CustomLB"

sudo docker pull $IMAGE
sudo docker stop $CONTAINER_NAME
sudo docker rm $CONTAINER_NAME
sudo docker run --name $CONTAINER_NAME -p 8080:8080 -tdi $IMAGE

sudo docker exec $CONTAINER_NAME /bin/bash -c "$PROJECT_PATH/main" &