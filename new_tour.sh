#!/bin/env bash
# new tour helper script
# usages:
#    bash new_tourh.sh for flowcontrol/1

TOUR_NAME=$1
RELAVITE_LINK=$2

mkdir "${TOUR_NAME}"

echo "/* ****************************************************************************
 * Tour: ${TOUR_NAME}
 * Link: https://tour.golang.org/${RELAVITE_LINK}
 * ----------------------------------------------------------------------------
 * 
 * ***************************************************************************/

package main

import \"fmt\"

func main() {
}" > "${TOUR_NAME}/${TOUR_NAME}.go"