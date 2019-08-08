#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

# exit script on error
set -e

cd "$(dirname "$0")"

# import colors
source ./colors.sh

# move to project root dir from ./scripts to ./
cd ..

function readMostRecentTag(){
	local last=$(git describe --tags)
	echo "$last"
}

function executeBuild(){
	# 1 store new tag value in VERSION file
	log "building new tag: $NEW_TAG"
	log "overwriting value of VERSION file with new tag content"
	echo $NEW_TAG > VERSION
	# verify file content
	fileContent=$(cat VERSION)
	if [[ $fileContent == $NEW_TAG ]]; then
		ok "VERSION file successfully updated"
	else
		fail "failed to updated VERSION file. aborting"
		exit
	fi

	# create requested tag
	log "creating tag $NEW_TAG"
	git tag -a $NEW_TAG -m "$NEW_TAG" 

	log "uploading current tag ($NEW_TAG) version"
	git push --tags

	log "creating current version docker images"

	log "creating current oss docker image"
	bash ./scripts/docker/docker_build_oss.sh

	log "creating current etherniti.org arm proxy docker image"
	bash ./scripts/docker/docker_build_pre_armhf.sh

	log "uploading etherniti/proy:latest image to dockerhub"
	docker push etherniti/proxy:latest 

	log "checking differences between last tag and current tag"
	diff=$(git log --pretty=oneline $PREVIOUS_TAG...$NEW_TAG)
	echo $diff
}

function cancelBuild(){
	fail "build aborted"
	exit
}

function main(){
	
	title "building next etherniti tagged version release"
	log "Reading last tagged value"
	PREVIOUS_TAG=$(readMostRecentTag)
	log "Last tagged value is: $PREVIOUS_TAG"


	input "Enter new tag version (current $PREVIOUS_TAG): "
	read NEW_TAG

	log "New tag value will be set to: $NEW_TAG."
	while true; do
	    read -p "Continue (yes/no): " yn
	    case $yn in
	        [Yy]* ) executeBuild NEW_TAG; break;;
	        [Nn]* ) cancelBuild; break;;
	        * ) echo "Please answer yes or no.";;
	    esac
	done
}

main
