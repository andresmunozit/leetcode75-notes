#!/bin/bash

# Check if jq is installed
if ! command -v jq &> /dev/null
then
    echo "jq could not be found. Please install it."
    exit 1
fi

# Check if the correct number of arguments is given
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 path/to/package.json"
    exit 1
fi

# File paths
PACKAGE_JSON="$1"
SCRIPTS_JSON="leetcode-node-scripts.json"

# Check if the package.json file exists
if [ ! -f "$PACKAGE_JSON" ]; then
    echo "package.json not found at $PACKAGE_JSON"
    exit 1
fi

# Check if the leetcode-node-scripts.json file exists
if [ ! -f "$SCRIPTS_JSON" ]; then
    echo "$SCRIPTS_JSON not found"
    exit 1
fi

# Replace the scripts section in package.json
jq --argfile scripts "$SCRIPTS_JSON" '.scripts = $scripts' "$PACKAGE_JSON" > temp.json && mv temp.json "$PACKAGE_JSON"

echo "Scripts section has been updated in $PACKAGE_JSON"
