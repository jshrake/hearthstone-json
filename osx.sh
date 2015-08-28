#!/usr/bin/env bash
# This script will acquire the hearthstone-json binary for your platform
# and extract
set -e
# set -v

# Disunity requires java! Sorry!
if ! type -p java >> /dev/null; then
  echo "Error: this script requires java for disunity, a program to extract assets from unity archive files";
  exit 1;
fi

# Optionally take path to cardxml0.unity3d
PATH_TO_CARDDEF=$1
if [ -z $PATH_TO_CARDDEF ];
  then
  PATH_TO_CARDDEF="/Applications/Hearthstone/Data/OSX/cardxml0.unity3d"
fi

if [ ! -f ${PATH_TO_CARDDEF} ]
  then
  echo "File not found: ${PATH_TO_CARDDEF}"
  echo "Please provide a path to your Hearthstone cardxml0.unity3d file"
  exit 1
fi

# Download the appropriate hearthstone-json binary
ARCH=`uname -m`
HEARTHSTONE_JSON="hearthstone-json-${ARCH}-macosx"
# TODO(jshrake): Acquire binary from Release binaries

# Use disunity to extract the card def xml files and transform each to json
DISUNITY_VERS=0.3.4
curl -LOk https://github.com/ata4/disunity/releases/download/v${DISUNITY_VERS}/disunity_v${DISUNITY_VERS}.zip
tar -xzvf disunity_v${DISUNITY_VERS}.zip disunity.jar lib
cp ${PATH_TO_CARDDEF} cardxml0.unity3d
java -jar disunity.jar extract cardxml0.unity3d
rm cardxml0.unity3d

# Transform the xml files
mkdir -p output
for file in cardxml0/CAB-cardxml0/TextAsset/*.txt; do
  filename=$(basename "$file")
  extension="${filename##*.}"
  filename="${filename%.*}"
  output="output/${filename}.json"
  echo "Generating ${output} from ${file}";
  ./hearthstone-json ${file} > ${output};
done
echo "Finished! Please find the JSON files in the output directory"

# Clean up
rm -rf disunity*
rm -rf cardxml0*
rm -rf lib