package main

import (
	"fmt"
	"go/helper"
	"sync"
	"time"
)

const conferencetickets int = 50

var conferencename = "Go conference"
var Remainingtickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstname       string
	lastname        string
	email           string
	numberoftickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetusers()

	firstname, lastname, email, usertickets := getuserinput()
	isvalidname, isvalidemail, isvalidticketnumber := helper.Validateuserinput(firstname, lastname, email, usertickets, Remainingtickets)

	if isvalidname && isvalidemail && isvalidticketnumber {

		BookTicket(usertickets, firstname, lastname, email)

		wg.Add(1)
		go sendticket(usertickets, firstname, lastname, email)

		firstname := getFirstNames()
		fmt.Printf("the first name of our bookings are %v\n", firstname)

		if Remainingtickets == 0 {
			//end program
			fmt.Printf("Sorry %v tickets are booked out.come back next year \n", conferencename)
			//break
		}
	} else {
		if !isvalidname {

			fmt.Println("first name or last name you entered is too short")
		}
		if !isvalidemail {
			fmt.Println("email address you entered dosen't contain @ sing")

		}
		if !isvalidticketnumber {
			fmt.Println("number of tickets you entered is invalid")
		}

	}
	wg.Wait()
}

func greetusers() {
	fmt.Printf("welcome to our %v\n", conferencename)
	fmt.Printf("In this %v you will know everything about Go\n", conferencename)
	fmt.Printf("If you want to attend this %v .you need to buy the exclusive  %v tickets to attend this grand event.\n ", conferencename, conferencename)
	fmt.Printf("We have total of %v and %v are still remaining\n", conferencetickets, Remainingtickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstname)
	}
	return firstNames

}

func getuserinput() (string, string, string, uint) {
	var firstname string
	var lastname string
	var email string
	var usertickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstname)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastname)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&usertickets)

	return firstname, lastname, email, usertickets
}

func BookTicket(usertickets uint, firstname string, lastname string, email string) {
	Remainingtickets = Remainingtickets - usertickets

	// create a map for a user
	var userdata = UserData{
		firstname:       firstname,
		lastname:        lastname,
		email:           email,
		numberoftickets: usertickets,
	}

	bookings = append(bookings, userdata)
	fmt.Printf("list of bookings is %v\n", bookings)

	fmt.Printf("thank you %v %v for booking %v tickets. You will recive a confirmation email at %v\n ", firstname, lastname, usertickets, email)
	fmt.Printf("still %v tickets remaining for %v\n", Remainingtickets, conferencename)

}
func sendticket(usertickets uint, firstname string, lastname string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", firstname, lastname, usertickets)
	fmt.Println("##############")
	fmt.Printf("sending ticket:\n %v\nto email address %v\n", ticket, email)
	fmt.Println("##############")
	wg.Done()
}
