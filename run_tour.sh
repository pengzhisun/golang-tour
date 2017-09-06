#!/bin/env bash

TOUR_NAME=$1
PARENT_DIR=${2:-.}

pushd . > /dev/null
cd ${PARENT_DIR}/${TOUR_NAME} > /dev/null
go build > /dev/null
./${TOUR_NAME}
popd > /dev/null