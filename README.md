# hearthstone-json

A command-line tool for converting the bundled [Hearthstone](http://us.battle.net/hearthstone/) card definitions from XML to JSON.

## What
---

Hearthstone ships with a [unity3d](https://unity3d.com/) asset bundle containing XML files for every supported language. The XML files contain definitions for every card in the game. The default install location for the unity3d asset bundle is:

- Windows: C:\Program Files (x86)\Hearthstone\Data\Win\cardxml0.unity3d
- OSX: /Applications/Hearthstone/Data/OSX/cardxml0.unity3d

In order to extract the XML files from the bundle, we need to use a tool such as [disunity](https://github.com/ata4/disunity), which depends on Java. The purpose of this tool is to transform the extracted XML files to human-readable JSON. The resulting JSON is compatible with the API defined in [http://hearthstonejson.com/](http://hearthstonejson.com/).

## Usage
---

### OSX
```bash
curl https://raw.githubusercontent.com/jshrake/hearthstone-json/master/osx.sh | sh
```

The above performs the heavy lifting discussed in [what](#what), specifically:

- Checks that java is installed and on your PATH
- Downloads the `hearthstone-json` binary for your platform and architecture
- Downloads [disunity](https://github.com/ata4/disunity)
- Extracts the XML files from the `cardxml0.unity3d` file using disunity
- Runs the `hearthstone-json` binary on each of the XML files
- Writes the resulting JSON output to the `output` directory


## Developement
---

```bash
git clone https://github.com/jshrake/hearthstone-json
cd hearthstone-json
go install
```

Assuming you've extracted the xml content from the Hearthstone game client cardxml0.unity3d file:

```bash
hearthstone-json path/to/enUS.txt
```
