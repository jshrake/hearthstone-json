#!/usr/bin/env bash
set -e

echo "\"HELLO HELLO HELLO\" - Annoy-o-Tron"

# Disunity requires java! Sorry!
if ! type -p java >> /dev/null; then
  echo "Error: this script uses disunity, which requires java -- is it on your path?";
  exit 1
fi


# Optionally take path to cardxml0.unity3d
PATH_TO_CARDDEF=$1
if [ -z "$PATH_TO_CARDDEF" ]; then
  PATH_TO_CARDDEF="/Applications/Hearthstone/Data/OSX/cardxml0.unity3d"
fi

if [ ! -f $PATH_TO_CARDDEF ]; then
  echo "Error: File not found $PATH_TO_CARDDEF"
  echo "Please provide a path to your Hearthstone cardxml0.unity3d file"
  exit 1
fi

# Download the appropriate hearthstone-json binary
ARCH=$(uname -m)
HEARTHSTONE_JSON="hearthstone-json-${ARCH}-macosx"
HEARTHSTONE_JSON_VERSION="v0.1.0-alpha"
ACQUIRE_HEARTHSTONE_JSON="curl -LOk https://github.com/jshrake/hearthstone-json/releases/download/${HEARTHSTONE_JSON_VERSION}/${HEARTHSTONE_JSON}"
echo "$ACQUIRE_HEARTHSTONE_JSON"
$ACQUIRE_HEARTHSTONE_JSON
chmod +x "$HEARTHSTONE_JSON"

# Use disunity to extract the card def xml files and transform each to json
DISUNITY_VERSION=0.3.4
ACQUIRE_DISUNITY="curl -LOk https://github.com/ata4/disunity/releases/download/v${DISUNITY_VERSION}/disunity_v${DISUNITY_VERSION}.zip"
echo "$ACQUIRE_DISUNITY"
$ACQUIRE_DISUNITY
tar -xzvf "disunity_v${DISUNITY_VERSION}.zip" disunity.jar lib
cp "$PATH_TO_CARDDEF" cardxml0.unity3d
java -jar disunity.jar extract cardxml0.unity3d
rm cardxml0.unity3d

# Transform the xml files
for file in cardxml0/CAB-cardxml0/TextAsset/*.txt; do
  filename=$(basename "$file")
  filename="${filename%.*}"
  output="${filename}.json"
  echo "Generating $output from $file"
  ./"$HEARTHSTONE_JSON" "$file" > "$output"
done

# Clean up
rm -rf disunity*
rm -rf cardxml0*
rm -rf lib
rm "$HEARTHSTONE_JSON"

echo "\"Hellooo....\" - Annoy-o-Tron"
