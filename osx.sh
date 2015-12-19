#!/usr/bin/env bash
set -e

echo "\"HELLO HELLO HELLO\" - Annoy-o-Tron"

# Download the hearthstone datafiles
HEARTHSTONE_DATA="https://github.com/jshrake/extracted-hearthstone-data/archive/master.zip"
ACQUIRE_HEARTHSTONE_DATA="curl -LO ${HEARTHSTONE_DATA}"
echo "$ACQUIRE_HEARTHSTONE_DATA"
$ACQUIRE_HEARTHSTONE_DATA
HEARTHSTONE_DATA_DIR="./extracted-hearthstone-data-master"
rm -rf $HEARTHSTONE_DATA_DIR
unzip master.zip

# Download the appropriate hearthstone-json binary
ARCH=$(uname -m)
HEARTHSTONE_JSON="hearthstone-json-${ARCH}-macosx"
HEARTHSTONE_JSON_VERSION="v0.1.1-alpha"
ACQUIRE_HEARTHSTONE_JSON="curl -LOk https://github.com/jshrake/hearthstone-json/releases/download/${HEARTHSTONE_JSON_VERSION}/${HEARTHSTONE_JSON}"
echo "$ACQUIRE_HEARTHSTONE_JSON"
$ACQUIRE_HEARTHSTONE_JSON
chmod +x "$HEARTHSTONE_JSON"

# Transform the xml files
for file in ${HEARTHSTONE_DATA_DIR}/cardxml0/*.txt; do
  filename=$(basename "$file")
  filename="${filename%.*}"
  output="${filename}.json"
  echo "Generating $output from $file"
  ./"$HEARTHSTONE_JSON" "$file" > "$output"
done

# Clean up
rm -rf ${HEARTHSTONE_DATA_DIR}
rm master.zip
rm "$HEARTHSTONE_JSON"

echo "\"Hellooo....\" - Annoy-o-Tron"
