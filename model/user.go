package model

import (
	"fmt"
	"strings"
)

//Users struct for Users.
type Users struct {
	List []User `json:"users"`
}

//User struct for User.
type User struct {
	Logname           string   `json:"logname"`
	Type              string   `json:"type"`
	UID               int      `json:"uid"`
	FullName          string   `json:"fullname"`
	InitialPassword   string   `json:"initial_password"`
	RequireAnnotation bool     `json:"require_annotation"`
	Roles             []string `json:"roles"`
	KioskMode         bool     `json:"kiosk_mode"`
	KioskScreen       string   `json:"kiosk_screen"`
	HREF              string   `json:"href"`
}

//PrintUserInfo to print chassis component info.
func (user *User) PrintUserInfo() {
	line := fmt.Sprintln("---------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("%-15s %-10s %-11d %-35s %-8v\n%s", user.Logname, user.Type, user.UID, strings.Join(user.Roles, ","), user.RequireAnnotation, line)
}
