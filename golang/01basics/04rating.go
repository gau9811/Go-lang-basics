package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var userRating string

	//Front end
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your full name")
	name, _ = reader.ReadString('\n')

	// take rating from user and convert it to float
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("please rate our burger center between 1 and 5:")
	userRating, _ = reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	//Back end
	fmt.Printf("Hello %v, \n Thanks for rating our burger center with %v star rating. \n\n Your rating was recorded in our system at %v", name, mynumRating, time.Now().Format(time.Kitchen))

	if mynumRating == 5 {	
		fmt.Println("\n Bonus for team for 5 star service")
	} else if mynumRating == 4 || mynumRating == 3 {
		fmt.Println("\n we are imporoving")
	} else if mynumRating <= 3 {
		fmt.Println("\n we need to do series work on burger")
	}

}
