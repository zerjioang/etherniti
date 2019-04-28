#!/bin/bash

echo "Generating new mkdocs project"
docker exec cc_mkdocs_container mkdocs new etherniti_docs
if [[ -d data/etherniti_docs ]]; then
	echo "Chaincode documentation folder successfully created"
	ls -alh data/etherniti_docs
	sudo chown $USER -R data
	echo ""
	echo "Edit the files and once you finish, execute generate_docs.sh"
	echo ""
fi