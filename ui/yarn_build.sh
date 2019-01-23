#!/bin/bash

# move to project root dir
cd "$(dirname "$0")"

COOMING_SOON=true yarn build
cp -ra robots.txt ./docs/robots.txt


# check the extension of index file and
# force to be html
if [[ -f index.template ]]; then
	mv index.template index.html
fi