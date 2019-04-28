#!/bin/bash

echo "serving documentation on 0.0.0.0:8888"
docker exec cc_mkdocs_container sh -c "cd etherniti_docs && mkdocs serve --dev-addr 0.0.0.0:8888"