#!/bin/bash

PLUGIN_NAME=$1
CODEOWNERS="${@:2}" # GitHub IDs of each codeowner split by ' '. They must be started by '@'.

if [ -z "$PLUGIN_NAME" -o -z "$CODEOWNERS" ]; then
    echo "Usage: $0 <plugin-name> <codeowner1> <codeowner2> ..."
    exit 1
fi

# 1. Initialize the plugin directory.
PLUGIN_DIR=plugins/$PLUGIN_NAME
mkdir $PLUGIN_DIR

# 1-1. README
cp hack/init-template/README.md $PLUGIN_DIR/README.md
# Replace the issues link
gsed -i "s|{{ISSUES_PLUGIN_NAME}}|${PLUGIN_NAME}|g" $PLUGIN_DIR/README.md
# Replace the codeowners link
codeowner_links=""
for codeowner in $CODEOWNERS; do
    codeowner_links="$codeowner_links [@$codeowner](https://github.com/$codeowner) "
done
gsed -i "s|@{ACCOUNT}|$codeowner_links|" $PLUGIN_DIR/README.md

## 1-2. Makefile
cp hack/init-template/Makefile $PLUGIN_DIR/Makefile

# 1-3. go.mod
pushd $PLUGIN_DIR
go mod init github.com/pipe-cd/community-plugins/$PLUGIN_DIR
popd

# 2. Update Issue templates, CODEOWNERS
make sync
