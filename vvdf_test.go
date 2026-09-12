package vvdf

import (
	"fmt"
	"testing"
)

var testData = []string{
		`"AddonInfo"
{
     addonSteamAppID         "550"
     addontitle              "Super Healing" 
     addonversion            "2.0"
     addonauthor             "Andree Tran"
     addonDescription        "Super Healing" 
1
}`,
	`"AddonInfo"
{
     addonSteamAppID 			550			// 550 is the app ID for Left 4 Dead 2
     addontitle 			"Urban Flight"
     addonversion 			"12"
     addontagline 			"No plane, no gain"	// short description
     addonauthor			"The Rabbit"
     addonSteamGroupName		"group"
     addonauthorSteamID			"The_Rabbit42"
     addonContent_Campaign		1			// campaign mode included
     addonContent_Survival		1			// survival mode included
     addonContent_Scavenge		1			// scavenge mode included
     addonContent_Versus		1			// versus   mode included
     addonURL0				"http://steamcommunity.com/sharedfiles/filedetails/?id=121086524"
								// where people can download your VPK

     addonDescription			"The city is burning. As ash falls, the survivors attempt to flee to a small military airfield on the other side of the river. It's a straight path down the main boulevard, but nothing is ever simple. Crashed cars, blazing fires, police barricades, and the city itself all stand in their way."

     addonContent_Script		1			// Has Scripts
     addonContent_Music			0			// Has Custom Music
     addonContent_Sound			1			// Has Custom Sound
     addonContent_prop			1			// This Add-on provides new props,
     addonContent_Prefab		0			// Provides new prefabs
     addonContent_BackgroundMovie	0			// Provides a replacement for the background movie.
     addonContent_Survivor 		0			// Provides a new survivor model. 0=false, 1=true, String in quotes if replaces specific single character, i.e. "Coach"
     								// eg addonContent_Survivor "rochelle" works fine, no number.

     addonContent_BossInfected		1			// Provides a new boss infected model. Break these out?
     addonContent_CommonInfected	0			// Provides a new common infected model
     Content_WeaponModel		0			// Provides a new appearance to existing weapons, but does not change their function
     Content_weapon			0			// provides new weapons or new zombie killing functionality, i.e. guns, explosives, booby traps, hot tar,
     addonContent_Skin			1			// 0 if no new skin textures for existing models. 1 if multiple skin pack. String in quotes if specific single skin
     addonContent_Spray			0			// Provides new sprays.
     addonContent_Map			0			// Add-on provides a standalone map
}`,
	`missions:"mission"
{
	"Name"			"urbanflight"
	"Version"		"12"
	"Author"		"The Rabbit"
	"Website"		"http://steamcommunity.com/sharedfiles/filedetails/?id=121086524"

	"DisplayTitle"		"Urban Flight"
	"Description"		"The city is burning. As ash falls, the survivors attempt to flee to a small military airfield on the other side of the river. It's a straight path down the main boulevard, but nothing is ever simple. Crashed cars, blazing fires, police barricades, and the city itself all stand in their way."
	"Image"			"maps/uf0_poster"
	"OuttroImage"		"vgui/OutroTitle_UrbanFlight"

	"x360ctx"		"5"

	"no_wpn_restore"	"1"		// on player wipe, don't keep any weapons
	"meleeweapons"		"machete;fireaxe;frying_pan;cricket_bat;baseball_bat;knife;crowbar;golfclub;electric_guitar;katana;tonfa"

	// Loading poster data
	"poster"
	{
		"posterImage"		        "LoadingScreen_UrbanFlight"
		"posterImage_widescreen"	"LoadingScreen_UrbanFlight_widescreen"

		"fullscreen"			"1"

		"posterTitle"			"Urban Flight"
		"posterTitle_y"			"320"

		"posterTagline"			"No plane, no gain"
		"posterTagline_y"		"380"

		"l4d2_names"			"1"

		"mechanic_player_name_x"	"9999"
		"mechanic_player_name_y"	"9999"

		"coach_player_name_x"		"9999"
		"coach_player_name_y"		"9999"

		"producer_player_name_x"	"9999"
		"producer_player_name_y"	"9999"

		"gambler_player_name_x"		"9999"
		"gambler_player_name_y"		"9999"

		"character_order"		"mechanic;coach;producer;gambler"
	}

	"modes"
	{
		"coop"
		{
			"1"
			{
				"Map" "uf1_boulevard"
				"DisplayName" "1: Boulevard"
				"Image" "maps/uf1_boulevard"
			}
			"2"
			{
				"Map" "uf2_rooftops"
				"DisplayName" "2: Rooftops"
				"Image" "maps/uf2_rooftops"
			}
			"3"
			{
				"Map" "uf3_harbor"
				"DisplayName" "3: Harbor"
				"Image" "maps/uf3_harbor"
			}
			"4"
			{
				"Map" "uf4_airfield"
				"DisplayName" "4: Airfield"
				"Image" "maps/uf4_airfield"
			}
		}
		"versus"
		{
			"1"
			{
				"Map" "uf1_boulevard"
				"DisplayName" "1: Boulevard (VS)"
				"Image" "maps/uf1_boulevard"
				"VersusCompletionScore"	"500"
			}
			"2"
			{
				"Map" "uf2_rooftops"
				"DisplayName" "2: Rooftops (VS)"
				"Image" "maps/uf2_rooftops"
				"VersusCompletionScore"	"600"
				"versus_boss_spawning"
				{
					"spawn_pos_min"		"0.0"
					"spawn_pos_max"		"0.0"
					"tank_chance"		"0"
					"witch_chance"		"0"
					"witch_and_tank"	"0"
				}
			}
			"3"
			{
				"Map" "uf3_harbor"
				"DisplayName" "3: Harbor (VS)"
				"Image" "maps/uf3_harbor"
				"VersusCompletionScore"	"700"
			}
			"4"
			{
				"Map" "uf4_airfield"
				"DisplayName" "4: Airfield (VS)"
				"Image" "maps/uf4_airfield"
				"VersusCompletionScore"	"800"
			}
		}
		"survival"
		{
			"1"
			{
				"Map" "uf2_rooftops"
				"DisplayName" "Radio Station"
				"Image" "maps/uf2_station"
			}
			"2"
			{
				"Map" "uf4_airfield"
				"DisplayName" "Airfield"
				"Image" "maps/uf4_hanger"
			}
		}
		"scavenge"
		{
			"1"
			{
				"Map" "uf4_airfield"
				"DisplayName" "Airfield"
				"Image" "maps/uf4_backarea"
			}
		}
	}
}`,
	`"AddonInfo"
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

}`,
	`missions:// Mission files describe the metadata needed by campaign-specific add-ons so they can be
// integrated into Left4Dead. The data in this file is used by the game UI, matchmaking and server.
// Although you may provide multiple Campaigns in one add-on by putting more than one .TXT file
// in the missions folder, it's generally a good idea to stick to one per add-on.
//
// HOW TO DEBUG MISSION FILES:
//	In the console set "developer 2"
//	Then type "mission_reload"
// This will make the game reload all the mission files and print out every chapter for every mode it
// finds.  It's very useful to ensure that your mission file is being correctly read.

"mission"
{
	// Use a short name, since it is used as a tag on the servers for matching
	// your campaign when looking for a dedicated server.  Generally it should
	// be something unique.  One suggestion to is use your initials and a short
	// abbreviated name for your campaign. Avoid spaces and special characters.
	// Do not change the name when you create a revision, as the matchmaking
        // system will consider it a different campaign. Instead, use "Version" and
        // "DisplayTitle" below to indicate revisions.
	"Name"		"warcelona"

	// The version number is used when finding campaigns on dedicated servers
	// and in matchmaking. If you try to connect to someone in matchmaking
	// who has a newer version, you will be directed to download the new
	// version.  You must increment this to the next integer (whole numbers)every
        // time you release an update. (I.E. 1, 2, 3, 4, etc.)
	"Version"       "4"

	// Author is displayed in the UI when people see others playing your
	// campaign.
	"Author"	"Warcelona dev team"

	// Website is extremely important as people will automatically be sent
	// to this URL if they try to join a friend running your campaign.  This
	// should be the home page for your campaign and it should provide a
	// description, a link to download it, and installation instructions.
	"Website"	"http://www.warcelonacampaign.com/"

	// This name is used when referring to the campaign in the UI.
	"DisplayTitle"	"Warcelona"
	"Description"  	"Kill zombies in the streets of Barcelona. Football fans? Certainly. Beating zombies to death with a ham leg? That too. Spanish fiesta? Of course. Bulls? No doubt."

	// Vmt shown behind the end credits when your campaign completes.
        // Note: This item needs to have "vgui\" at the front. It is assumed
	// for the poster and the map thumbnails.
 	"OuttroImage"	"vgui\bdd_thumb_end"

	// Loading poster data
	//
	// Note that "posterTitle" and "posterTagline" are often blank as
	// the poster usually includes these in the poster image itself.
	// If you do not provide a poster, a generic default will be used.
	"poster"
	{
		"posterImage"		"cartel"
		"posterImage_widescreen"	"cartel"

		"fullscreen"		"1"

                //Note L4D2 does not position player names over the poster.

	}

	// The modes section lists each of your campaign maps and each mode
	// they support.  Depending on how you set up your campaign maps,
	// you can reuse the same bsp for different modes as is the case with
        // coop and versus in this example. This requires setting up entities
        // that show up or operate differently for different modes.
        //
	// The following modes are supported: "coop", "versus", "survival"
	//
	// Number each map section starting with "1"
	//
	// "Map" is the name of the bsp of the starting map. (do not include
        // an extension)
	// "DisplayName" is the name to use for the chapter of the map.
	// "Image" is the image used for the chapter in the lobby and
	// settings.

	"modes"
	{

		"coop"
		{
			"1"
			{
				"Map" "srocchurch"
				"DisplayName" "neighborhood"
				"Image" "maps/bdd_thumb1"
			}
			"2"
			{
				"Map" "plaza_espana"
				"DisplayName" "plaza"
				"Image" "maps/bdd_thumb2"
			}
			"3"
			{
				"Map" "maria_cristina"
				"DisplayName" "avenue"
				"Image" "maps/bdd_thumb3"
			}

			"4"
			{
				"Map" "mnac"
				"DisplayName" "gardens"
				"Image" "maps/bdd_thumb4"
				"SpawnBossThreats" "1"
				"coop_boss_spawning"
				{
					"spawn_witches"					"0"
					"spawn_tanks"					"1"
				}
			}

		}

	"versus"
		{
			"1"
			{
				"Map" "srocchurch"
				"DisplayName" "neighborhood"
				"Image" "maps/bdd_thumb1"
			}
			"2"
			{
				"Map" "plaza_espana"
				"DisplayName" "plaza"
				"Image" "maps/bdd_thumb2"
			}
			"3"
			{
				"Map" "maria_cristina"
				"DisplayName" "avenue"
				"Image" "maps/bdd_thumb3"
			}

			"4"
			{
				"Map" "mnac"
				"DisplayName" "gardens"
				"Image" "maps/bdd_thumb4"
				"VersusModifier" "1.5"
				"versus_boss_spawning"
				{
					"spawn_pos_min"		"0.71"
					"spawn_pos_max"		"0.75"
					"tank_chance"		"1"
					"witch_chance"		"0"
					"witch_and_tank"		"0"
				}
			}

		}

	}

}
`,
}

func TestStringToMap(t *testing.T) {
	for idx, el := range testData {
		tmap, err := StringToMap(el)

		fmt.Println("---------------", idx, "----------")
		if err != nil {
			panic("went wrong")
		}

		fmt.Printf("%+v\n\n", tmap)
	}
}
