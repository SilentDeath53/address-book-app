package main

import (
	"fmt"
	"os"
	"bufio"
)

type Contact struct {
	Name  string
	Phone string
}

func main() {
	contacts := make([]Contact, 0)

	for {
		fmt.Println("\nAddress Book")
		fmt.Println("1. Add Contact")
		fmt.Println("2. View Contacts")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addContact(&contacts)
		case 2:
			viewContacts(contacts)
		case 3:
			fmt.Println("Exiting the program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addContact(contacts *[]Contact) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter phone number: ")
	scanner.Scan()
	phone := scanner.Text()

	contact := Contact{
		Name:  name,
		Phone: phone,
	}

	*contacts = append(*contacts, contact)

	fmt.Println("Contact added successfully.")
}

func viewContacts(contacts []Contact) {
	fmt.Println("\nContacts:")
	for _, contact := range contacts {
		fmt.Printf("Name: %s, Phone: %s\n", contact.Name, contact.Phone)
	}
}
