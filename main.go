package main

import (
	"booking-app/common"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	emil        string
	noOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := common.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTickets(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookig are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference tickets are sold out, please come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address doesn't contain @ symbol")
			}
			if !isValidTicketNumber {
				fmt.Println("Ticket number is invalid")
			}
		}
	}

	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to conference %s booking application \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and tickets remaining are %v \n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend %s \n", conferenceName)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name : ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email : ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets : ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		emil:        email,
		noOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings : %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)
	fmt.Println("##########################")
	fmt.Printf("Sending ticket : \n %v \nto email address %v\n", ticket, email)
	fmt.Println("##########################")
	wg.Done()
}
