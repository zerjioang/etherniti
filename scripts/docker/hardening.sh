#!/bin/sh

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

set -x #trace on
set -e #break on error
echo "hardening container..."

# docker image hardening
sed -i -r "s/^$USER:!:/$USER:x:/" /etc/shadow

# remove unnecessary user accounts.
sed -i -r "/^($USER|root|appgroup)/!d" /etc/group
sed -i -r "/^($USER|root|appgroup)/!d" /etc/passwd
sed -i -r "/^($USER|root|appgroup)/!d" /etc/shadow

# reinforce access permissions
sysdirs="
/bin
/etc
/lib
/sbin
/usr
"
find $sysdirs -xdev -type d \
-exec chown root:root {} \; \
-exec chmod 0755 {} \;

