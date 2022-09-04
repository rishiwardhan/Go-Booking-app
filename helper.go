package helper

import "strings"

func Validateuserinput(firstname string, lastname string, email string, usertickets uint, remainingtickets uint) (bool, bool, bool) {
	isvalidname := len(firstname) >= 2 && len(lastname) >= 2
	isvalidemail := strings.Contains(email, "@")
	isvalidticketnumber := usertickets > 0 && usertickets <= remainingtickets
	return isvalidname, isvalidemail, isvalidticketnumber
}
