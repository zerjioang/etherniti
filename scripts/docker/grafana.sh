#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

# exit script on error
set -e

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ../..

echo "starting grafana service container"
# The default credentials are admin/admin.
# (You can of course change these after login.)
docker run -d \
    -p 3000:3000 \
    --name grafana \
    --mount src=grafana-storage,dst=/var/lib/grafana \
    -e "GF_INSTALL_PLUGINS=grafana-simple-json-datasource" \
    grafana/grafana