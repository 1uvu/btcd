#!/bin/bash

docker build . -t bitlog/btcd --build-arg ARCH=amd64
