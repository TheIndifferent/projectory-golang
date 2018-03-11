package projectory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var indicentsJsonHeredoc = []byte(`
[
  [
    "New market launch",
    "You have to switch",
    "to a different project",
    "until you reach",
    "100 points."
  ],
  [
    "The robber got in.",
    "Again.",
    "no calculators",
    "for 5 tasks."
  ],
  [
    "New EU legislation",
    "The company must be",
    "compliant. Read out",
    "loud following words:",
    "Privacy, Freedom,",
    "Compliance, Security,",
    "User rights, Policy,",
    "Change management,",
    "Improvement, Protocol,",
    "Regulations,",
    "Requirement."
  ],
  [
    "NSA needs [censored]",
    "[censored] [censored]",
    "Ask every member of",
    "the project to write",
    "down their first and",
    "last name on a paper,",
    "in alphabetical order."
  ],
  [
    "VPN failure",
    "You have to calculate",
    "next 2 cards before",
    "writing down",
    "the result."
  ],
  [
    "VAB",
    "You are blocked",
    "for 30 seconds."
  ],
  [
    "Company bets changed",
    "You have to embedd to",
    "other project. Switch",
    "to different project for 3",
    "tasks."
  ],
  [
    "Goalie duties",
    "You have to take over",
    "the project results",
    "writedown."
  ],
  [
    "Chapter commitment",
    "Count the dots:"
  ],
  [
    "Coaching a squad mate",
    "You have to read the",
    "next card to your",
    "squad mate out loud",
    "and (s)he has to solve it."
  ],
  [
    "Chrismas holidays",
    "Only one person per",
    "project can pick up",
    "cards for 30 seconds."
  ],
  [
    "Sick leave",
    "You are blocked",
    "for 30 seconds."
  ],
  [
    "PO is on vacation",
    "When you pick next card",
    "put it below the deck",
    "until you get a",
    "division operation."
  ],
  [
    "Moving to the new office",
    "Change seating",
    "arrangements so that",
    "no one sits on the",
    "same place."
  ],
  [
    "TAP",
    "Count the number of",
    "finished cards",
    "out loud."
  ],
  [
    "Backlog grooming",
    "Go through the remaining",
    "cards and count number",
    "of multiplication and",
    "division operations,",
    "then tell that number to",
    "the whole squad.",
    "Shuffle the deck."
  ]
]
`)

func ReadOrWriteIncidents(wd string) [][]string {
	filename := path.Join(wd, "incidents.json")
	_, statErr := os.Stat(filename)
	if statErr != nil {
		if os.IsNotExist(statErr) {
			LOG.Print("creating new 'incidents.json' file... ")
			ioutil.WriteFile(filename, indicentsJsonHeredoc, 0644)
		} else {
			LOG.Fatalln("Error checking for existing 'incidents.json' file")
		}
	}
	LOG.Print("reading 'incidents.json' file... ")
	data, readErr := ioutil.ReadFile(filename)
	if readErr != nil {
		LOG.Fatalln("Error reading 'incidents.json' file")
	}
	LOG.Print("parsing incidents... ")
	var incidents [][]string
	jsonErr := json.Unmarshal(data, &incidents)
	if jsonErr != nil {
		LOG.Fatalln("Reading JSON failed: " + fmt.Sprintf("%v", jsonErr))
	}
	LOG.Println("number of incidents: " + fmt.Sprintf("%d", len(incidents)))
	return incidents
}
