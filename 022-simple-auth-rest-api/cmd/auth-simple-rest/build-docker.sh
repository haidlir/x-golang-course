#!/bin/sh
echo "->>> Compiling into Alpine's binary file..."
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./app -v ./app.go
if [ $? -ne 0 ]; then
    echo "->>> Compiling Failed."
    exit -1
fi
echo "->>> Compiling Success."
echo "->>> Buidling into docker image..."
docker build -t haidlir/arkedemy-golang-course-app:latest -f Dockerfile .
if [ $? -ne 0 ]; then
    echo "->>> Building Failed."
    exit -1
fi
echo "->>> Deleting Binary File..."
rm ./app
echo "->>> Building Success."