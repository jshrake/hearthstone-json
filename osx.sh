#!/usr/bin/env bash

set -e

if ! type -p java >> /dev/null; then
  echo "Cannot find a java installation on this computer"
  echo "This script uses disunity, a java program to extract the XML card definitions";
  exit 1;
fi

# Optionally take path to cardxml0.unity3d
PATH_TO_CARDDEF=$1
if [ -z $PATH_TO_CARDDEF ];
  then
  PATH_TO_CARDDEF="/Applications/Hearthstone/Data/OSX/cardxml0.unity3d"
fi
echo "Path to cardxml0.unity3d file: ${PATH_TO_CARDDEF}"
cp ${PATH_TO_CARDDEF} carddefs.unity3d

# Use disunity to extract the card def xml files and transform each to json
# Acquire disunity
DISUNITY_VERS=0.3.4
echo "Acquiring disunity_v${DISUNITY_VERS}"
curl -LOk https://github.com/ata4/disunity/releases/download/v${DISUNITY_VERS}/disunity_v${DISUNITY_VERS}.zip
tar -xzvf disunity_v${DISUNITY_VERS}.zip disunity.jar lib
java -jar disunity.jar extract carddefs.unity3d
# Transform the xml files
mkdir -p output
for file in carddefs/CAB-cardxml0/TextAsset/*.txt; do
  filename=$(basename "$file")
  extension="${filename##*.}"
  filename="${filename%.*}"
  echo "Converting $file to output/${filename}.json";
  ./hearthstone-json ${file} > output/${filename}.json;
done
# Clean up
rm -rf disunity*
rm -rf carddefs*
rm -rf lib