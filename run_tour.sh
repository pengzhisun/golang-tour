#!/bin/env bash

TOUR_NAME=$1

pushd . > /dev/null
cd ${TOUR_NAME} > /dev/null
go build > /dev/null
./${TOUR_NAME}
popd > /dev/null