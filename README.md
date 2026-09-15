## vvdf

`vvdf is a lightweight, zero-dependency Go library for parsing Valve Data Format (VDF / KeyValues) text into Go maps. It provides a robust, single-pass parser designed to handle complex nested structures, inline comments, and non-standard Valve configuration formats safely.

> **Note on Test Coverage / Scope:**  
> *This library is tailored for VPK metadata processing and has been tested specifically against internal VPK configuration files: `addoninfo.txt` and `mission/*.txt`.*

Features
- Single-Pass Parser: Memory-efficient stream parsing using a bottom-up assembly strategy.

- Comment Awareness: Strips // comments cleanly while preserving URLs and quoted text.

- Zero Dependencies: Uses only the Go standard library (bufio, strings, regexp).

- Source Engine Compatible: Robust against irregular whitespace, unquoted tokens, and deeply nested block scopes.


## Installation

go get [github.com/bahyqn/vvdf](https://github.com/bahyqn/vvdf)
```
go get https://github.com/bahyqn/vvdf
```


## Usage

```
import "github.com/bahyqn/vvdf"

func main(){
    text := ""

    tmap, err := vvdf.StringToMap(text)
}
```

## Handled Edge Cases

1. Case 1 (`comments anywhere`)
```
"AddonInfo"
{
     addonSteamAppID         550                                                     	// 500 is the app ID for Left 4 Dead, 550 for Left 4 Dead 2
     addontitle              "Warcelona"                                          // Add-on title that shows up in Add-ons list. ~20 chars max
     addonversion	     1.4                                                     	// Add-on version.
     addontagline            "Kill zombies in the streets of Barcelona."	     	// Add-on tagline or wrap-up- a short description. ~100 chars max
     addonauthor             "Warcelona dev team"                                               // Name/alias of the author
     addonContent_Campaign   1 	                                                     	//This addon provides muliple connected maps with a finale
     addonURL0               "http://www.warcelonacampaign.com/" 		//An html home page for the add-on that includes a download link.

     // short description that appears in the Add-on list screen...
     addonDescription        "Kill zombies in the streets of Barcelona. Football fans? Certainly. Beating zombies to death with a ham leg? That too. Spanish fiesta? Of course. Bulls? No doubt."

addonContent_Script 1
addonContent_Music 1
addonContent_Sound 1
addonContent_prop 1 //This Add-on provides new props,
addonContent_Prefab 0 //Provides new prefabs
addonContent_BossInfected 1

addonContent_Skin 1 //0 if no new skin textures for existing models. 1 if multiple skin pack. String in quotes if specific single skin

 
}
```

2. Case 2 (`dirty letter,no "`)
```
"AddonInfo"
{
     addonSteamAppID         "550"
     addontitle              "Super Healing" 
     addonversion            "2.0"
     addonauthor             "Andree Tran"
     addonDescription        "Super Healing" 
1 <===================================================================== here
}


`missions:"mission"
{
	"Name"			"urbanflight"

	"poster"
	{
		"posterImage"		        "LoadingScreen_UrbanFlight"
	}

	"modes"
	{
       1 <===================================================================== here
		"coop"
		{
			"1"
			{
				"Map" "uf1_boulevard"
				"DisplayName" "1: Boulevard"
				"Image" "maps/uf1_boulevard"
			}

		}
	}
}
```

3. Case 3
```
addonTitle              "Super Healing" 

addontitle              "Super Healing" 
```

4. Case 4 (`without }`)
```
"AddonInfo"
{
        "addonSteamAppID"                  "550"
        "addonTitle"                       "The Curse of Lazar Castle"
        "addonVersion"                     "1.0"
        "addonTagline"                     "God help us all"
        "addonAuthor"                      "Dereck"
        "addonAuthorSteamID"               "Kvothe Venture"
        "addonSteamGroupName"              "0"
        "addonURL0"                        "http://steamcommunity.com/sharedfiles/filedetails/?id=503361241"

< ========================== without }
```

5. Case 5
```

```