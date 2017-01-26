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

//FCTargets struct for FC targets.
type FCTargets struct {
	List []FCTarget `json:"targets"`
}

//FCTarget struct for FC target.
type FCTarget struct {
	DiscoveredPorts    int    `json:"discovered_ports"`
	HREF               string `json:"href"`
	InvalidCRCCount    int    `json:"invalid_crc_count"`
	InvalidTXWordCount int    `json:"invalid_tx_word_count"`
	LinkFailureCount   int    `json:"link_failure_count"`
	LossOfSignalCount  int    `json:"loss_of_signal_count"`
	LossOfSyncCount    int    `json:"loss_of_sync_count"`
	Mode               string `json:"mode"`
	Port               string `json:"port"`
	ProtocolErrorCount int    `json:"protocol_error_count"`
	Speed              string `json:"speed"`
	WWN                string `json:"wwn"`
}

//PrintTargetInfo to print target info.
func (target *FCTarget) PrintTargetInfo() {
	fmt.Printf("%-25s %-18s %-10s %-8s %-8d %-8d %-8d %8d\n", target.WWN, target.Port, target.Speed, target.Mode,
		target.DiscoveredPorts, target.LossOfSyncCount, target.LossOfSignalCount, target.LinkFailureCount)
}

//IscsiInitiators struct for Initiators.
type IscsiInitiators struct {
	List []IscsiInitiator `json:"initiators"`
}

//IscsiInitiator struct for Initiator.
type IscsiInitiator struct {
	Alias      string `json:"alias"`
	ChapSecret string `json:"chapsecret"`
	ChapUser   string `json:"chapuser"`
	HREF       string `json:"href"`
	Initiator  string `json:"initiator"`
}

//PrintInitiatorInfo to print initiator info.
func (initiator *IscsiInitiator) PrintInitiatorInfo() {
	fmt.Printf("%-55s %-15s %-11s %20s\n", initiator.Initiator, initiator.Alias, initiator.ChapUser, initiator.ChapSecret)
}

//IscsiInitiatorGroups struct for InitiatorsGroups.
type IscsiInitiatorGroups struct {
	List []IscsiInitiatorGroup `json:"groups"`
}

//IscsiInitiatorGroup struct for InitiatorGroup.
type IscsiInitiatorGroup struct {
	HREF       string   `json:"href"`
	Initiators []string `json:"initiators"`
	Name       string   `json:"name"`
}

//PrintInitiatorGroupInfo to print initiator info.
func (group *IscsiInitiatorGroup) PrintInitiatorGroupInfo() {
	line := fmt.Sprintln("---------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("name: %s\ninitiators: %s\n%s", group.Name, strings.Join(group.Initiators, " "), line)
}
