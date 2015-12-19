package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// XMl Structs
type Tag struct {
	EnumID      int    `xml:"enumID,attr"`
	Type        string `xml:"type,attr"`
	Value       int    `xml:"value,attr"`
	StringValue string `xml:",chardata"` //populated when Type=="string"
}

type Entity struct {
	Version string `xml:"version,attr"`
	CardID  string `xml:"CardID,attr"`
	Tags    []Tag  `xml:"Tag"`
}

type CardDefs struct {
	Entities []Entity `xml:"Entity"`
}

// Conversion tables, credit to https://github.com/Sembiance/hearthstonejson
var EnumIDToString = map[int]string{
	185: "CardName",
	183: "CardSet",
	202: "CardType",
	201: "Faction",
	199: "Class",
	203: "Rarity",
	48:  "Cost",
	251: "AttackVisualType",
	184: "CardTextInHand",
	47:  "Atk",
	45:  "Health",
	321: "Collectible",
	342: "ArtistName",
	351: "FlavorText",
	32:  "TriggerVisual",
	330: "EnchantmentBirthVisual",
	331: "EnchantmentIdleVisual",
	268: "DevState",
	365: "HowToGetThisGoldCard",
	190: "Taunt",
	364: "HowToGetThisCard",
	338: "OneTurnEffect",
	293: "Morph",
	208: "Freeze",
	252: "CardTextInPlay",
	325: "TargetingArrowText",
	189: "Windfury",
	218: "Battlecry",
	200: "Race",
	192: "Spellpower",
	187: "Durability",
	197: "Charge",
	362: "Aura",
	361: "HealTarget",
	349: "ImmuneToSpellpower",
	194: "Divine Shield",
	350: "AdjacentBuff",
	217: "Deathrattle",
	191: "Stealth",
	220: "Combo",
	339: "Silence",
	212: "Enrage",
	370: "AffectedBySpellPower",
	240: "Cant Be Damaged",
	114: "Elite",
	219: "Secret",
	363: "Poisonous",
	215: "Recall",
	340: "Counter",
	205: "Summoned",
	367: "AIMustPlay",
	335: "InvisibleDeathrattle",
	377: "UKNOWN_HasOnDrawEffect",
	388: "SparePart",
	389: "UNKNOWN_DuneMaulShaman",
	380: "UNKNOWN_Blackrock_Heroes",
	396: "UNKNOWN_Grand_Tournement_Fallen_Hero",
	401: "UNKNOWN_BroodAffliction",
	402: "UNKNOWN_Intense_Gaze",
	403: "Inspire",
	404: "UNKNOWN_Grand_Tournament_Arcane_Blast",
}

var CardSetIDToString = map[int]string{
	2:  "Basic",
	3:  "Classic",
	4:  "Reward",
	5:  "Missions",
	7:  "System",
	8:  "Debug",
	11: "Promotion",
	12: "Curse of Naxxramas",
	13: "Goblins vs Gnomes",
	14: "Blackrock Mountain",
	15: "The Grand Tournament",
	16: "Credits",
	17: "Hero Skins",
	18: "Tavern Brawl",
	20: "The League of Explorers",
}

var CardTypeIDToString = map[int]string{
	3:  "Hero",
	4:  "Minion",
	5:  "Spell",
	6:  "Enchantment",
	7:  "Weapon",
	10: "Hero Power",
}

var RarityIDToString = map[int]string{
	0: "undefined",
	1: "Common",
	2: "Free",
	3: "Rare",
	4: "Epic",
	5: "Legendary",
}

var CardRaceIDToString = map[int]string{
	14: "Murloc",
	15: "Demon",
	20: "Beast",
	21: "Totem",
	23: "Pirate",
	24: "Dragon",
	17: "Mech",
}

var ClassIDToString = map[int]string{
	0:  "undefined",
	2:  "Druid",
	3:  "Hunter",
	4:  "Mage",
	5:  "Paladin",
	6:  "Priest",
	7:  "Rogue",
	8:  "Shaman",
	9:  "Warlock",
	10: "Warrior",
	11: "Dream",
}

var FactionIDToString = map[int]string{
	1: "Horde",
	2: "Alliance",
	3: "Neutral",
}

var Mechanics = []string{
	"Windfury", "Combo", "Secret", "Battlecry", "Deathrattle",
	"Taunt", "Stealth", "Spellpower", "Enrage", "Freeze",
	"Charge", "Overload", "Divine Shield", "Silence", "Morph",
	"OneTurnEffect", "Poisonous", "Aura", "AdjacentBuff",
	"HealTarget", "GrantCharge", "ImmuneToSpellpower",
	"AffectedBySpellPower", "Summoned", "Inspire",
}

func IsMechanic(s string) bool {
	for _, mechanic := range Mechanics {
		if strings.ToLower(mechanic) == strings.ToLower(s) {
			return true
		}
	}
	return false
}

// JSON Card struct
type JsonCard struct {
	Name         string   `json:"name"`
	ID           string   `json:"id"`
	Attack       int      `json:"attack,omitempty"`
	Health       int      `json:"health,omitempty"`
	Cost         int      `json:"cost,omitempty"`
	Durability   int      `json:"durability,omitempty"`
	Type         string   `json:"type"`
	Rarity       string   `json:"rarity"`
	Text         string   `json:"text,omitempty"`
	InPlayText   string   `json:"inPlayText,omitempty"`
	Flavor       string   `json:"flavor,omitempty"`
	Set          string   `json:"set"`
	Class        string   `json:"playerClass,omitempty"`
	Race         string   `json:"race,omitempty"`
	Mechanics    []string `json:"mechanics,omitempty"`
	Faction      string   `json:"faction,omitempty"`
	Artist       string   `json:"artist,omitempty"`
	Collectible  bool     `json:"collectible"`
	Elite        bool     `json:"elite,omitempty"`
	HowToGet     string   `json:"howToGet,omitempty"`
	HowToGetGold string   `json:"howToGetGold,omitempty"`
}

func EntityToJson(e Entity) JsonCard {
	var card JsonCard
	card.ID = e.CardID
	for _, tag := range e.Tags {
		enum := EnumIDToString[tag.EnumID]
		switch {
		case enum == "CardName":
			card.Name = tag.StringValue
		case enum == "CardSet":
			card.Set = CardSetIDToString[tag.Value]
		case enum == "CardType":
			card.Type = CardTypeIDToString[tag.Value]
		case enum == "CardRace":
			card.Race = CardRaceIDToString[tag.Value]
		case enum == "Class":
			card.Class = ClassIDToString[tag.Value]
		case enum == "Rarity":
			card.Rarity = RarityIDToString[tag.Value]
		case enum == "Faction":
			card.Faction = FactionIDToString[tag.Value]
		case enum == "Atk":
			card.Attack = tag.Value
		case enum == "Health":
			card.Health = tag.Value
		case enum == "Durability":
			card.Durability = tag.Value
		case enum == "Cost":
			card.Cost = tag.Value
		case enum == "CardTextInHand":
			card.Text = tag.StringValue
		case enum == "CardTextInPlay":
			card.InPlayText = tag.StringValue
		case enum == "FlavorText":
			card.Flavor = tag.StringValue
		case enum == "Collectible":
			card.Collectible = tag.Value == 1
		case enum == "ArtistName":
			card.Artist = tag.StringValue
		case enum == "Elite":
			card.Elite = tag.Value == 1
		case enum == "HowToGetThisCard":
			card.HowToGet = tag.StringValue
		case enum == "HowToGetThisGoldCard":
			card.HowToGetGold = tag.StringValue
		case IsMechanic(enum):
			card.Mechanics = append(card.Mechanics, enum)
		}

	}
	if len(card.Class) == 0 {
		card.Class = "Neutral"
	}
	return card
}

func PrintUsageAndExit() {
	fmt.Println("Transform a Hearthstone xml file extracted from cardxml0.unity3d to JSON")
	fmt.Println("Usage: hearthstone-json path/to/enUS.txt")
	os.Exit(1)
}

func main() {
	// Get input from stdin or from an xml file supplied as the first
	// positional argument
	input := os.Stdin
	if len(os.Args) > 1 {
		arg := os.Args[1]
		// Handle help -- just in case!
		if arg == "-h" || arg == "--help" {
			PrintUsageAndExit()
		}
		// Open the xml file
		xmlPath, err := filepath.Abs(arg)
		if err != nil {
			fmt.Println(err)
			PrintUsageAndExit()
		}
		input, err = os.Open(xmlPath)
		if err != nil {
			fmt.Println(err)
			PrintUsageAndExit()
		}
	}
	defer input.Close()
	// Decode the XML
	var cardDefs CardDefs
	if err := xml.NewDecoder(input).Decode(&cardDefs); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Transform the data
	var jsonCards []JsonCard
	for _, entity := range cardDefs.Entities {
		jsonCards = append(jsonCards, EntityToJson(entity))
	}
	// Output JSON!
	jsonData, err := json.MarshalIndent(jsonCards, "", " ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	output := os.Stdout
	defer output.Close()
	output.Write(jsonData)
}
