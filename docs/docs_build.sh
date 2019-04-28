#!/bin/bash

echo "Generating production files..."
docker exec cc_mkdocs_container sh -c "cd etherniti_docs && mkdocs build"
if [[ -d data/etherniti_docs/site ]]; then
	echo "Chaincode documentation successfuly built"
	ls -alh data/etherniti_docs/site
	echo "Setting ownership of the docs..."
	sudo chown $USER -R data/etherniti_docs/site
	echo "Moving documentation to gitlab pages folder..."
	mv ./data/etherniti_docs/site ../
	mv ../site ../public
fi