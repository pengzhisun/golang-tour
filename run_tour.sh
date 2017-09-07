#!/bin/env bash
# run tour code helper script
# usages:
#   bash run_tour.sh for tour/flowcontrol

TOUR_NAME=$1
PARENT_DIR=${2:-.}

pushd . > /dev/null
cd ${PARENT_DIR}/${TOUR_NAME} > /dev/null
go build > /dev/null
echo "####################  Source  ####################"
cat "./${TOUR_NAME}.go"
echo "####################  Output  ####################"
./${TOUR_NAME}
popd > /dev/null