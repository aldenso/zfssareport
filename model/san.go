package model

import "fmt"
import "strings"

//FCInitiators struct for Initiators.
type FCInitiators struct {
	List []FCInitiator `json:"initiators"`
}

//FCInitiator struct for Initiator.
type FCInitiator struct {
	Alias     string `json:"alias"`
	HREF      string `json:"href"`
	Initiator string `json:"initiator"`
}

//PrintInitiatorInfo to print initiator info.
func (initiator *FCInitiator) PrintInitiatorInfo() {
	fmt.Printf("%-25s %-15s\n", initiator.Initiator, initiator.Alias)
}

//FCInitiatorGroups struct for InitiatorsGroups.
type FCInitiatorGroups struct {
	List []FCInitiatorGroup `json:"groups"`
}

//FCInitiatorGroup struct for InitiatorGroup.
type FCInitiatorGroup struct {
	HREF       string   `json:"href"`
	Initiators []string `json:"initiators"`
	Name       string   `json:"name"`
}

//PrintInitiatorGroupInfo to print initiator info.
func (group *FCInitiatorGroup) PrintInitiatorGroupInfo() {
	line := fmt.Sprintln("---------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("name: %s\ninitiators: %s\n%s", group.Name, strings.Join(group.Initiators, " "), line)
}
