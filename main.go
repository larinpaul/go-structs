//// 2023/10/14 // 14:16 //

// Structs in Go

// We use structs in Go to represent structured data. It's often convenient to
// group different types of variables together. For example, if we want to
// represent a car we could do the following:
// type car struct {
// 	Make   string
// 	Model  string
// 	Height int
// 	Width  int
// }

// This creates a new struct type called car.
// All cars have a Make, Model, Height and WIdth

// In Go, you will often use a struct to represent information that you would
// have used a dictionary for in Python, or an object literal for in JavaScript

// Assignment:
// Complete the messageToSend struct definition. It needs two fields:
// * phoneNumber - an integer
// * message - a string.

// package main

// import "fmt"

// type messageToSend struct {
// 	phoneNumber int
// 	message     string
// }

// func test(m messageToSend) {
// 	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
// 	fmt.Println("====================================")
// }

// func main() {
// 	test(messageToSend{
// 		phoneNumber: 1311334,
// 		message:     "Thanks for signing up",
// 	})
// 	test(messageToSend{
// 		phoneNumber: 315423321146251225,
// 		message:     "Love to have you aboard!",
// 	})
// 	test(messageToSend{
// 		phoneNumber: 5153136,
// 		message:     "We're so excited to have you",
// 	})
// }

//// 14:24
// Nested structs in Go

// type car struct {
// 	Make string
// 	Model string
// 	Height int
// 	Width int
// 	FrontWheel Wheel
// 	BackWheel Wheel
// }

// type Wheel struct {
// 	Radius int
// 	Material string
// }

// The fields of a struct can be accessed using the dot . operator.

// // This is how we can instantiate a new instance of a struct
// myCar := car{}
// myCar.FrontWheel.Radius = 5

// package main

// import (
// 	"fmt"
// )

// type messageToSend struct {
// 	message   string
// 	sender    user
// 	recipient user
// }

// type user struct {
// 	name   string
// 	number int
// }

// func canSendMessage(mToSend messageToSend) bool {
// 	if mToSend.sender.name == "" {
// 		return false
// 	}
// 	if mToSend.recipient.name == "" {
// 		return false
// 	}
// 	if mToSend.sender.number == 0 {
// 		return false
// 	}
// 	if mToSend.recipient.number == 0 {
// 		return false
// 	}
// 	return true
// }

// // don't touch below this line

// func test(mToSend messageToSend) {
// 	fmt.Printf(`sending "%s" from %s (%v) to %s (%v)...`,
// 		mToSend.message,
// 		mToSend.sender.name,
// 		mToSend.sender.number,
// 		mToSend.recipient.name,
// 		mToSend.recipient.number,
// 	)
// 	fmt.Println()
// 	if canSendMessage(mToSend) {
// 		fmt.Println("...sent!")
// 	} else {
// 		fmt.Println("...can't send message")
// 	}
// 	fmt.Println("====================================")
// }

// func main() {
// 	test(messageToSend{
// 		message: "you have an appointment tomorrow",
// 		sender: user{
// 			name:   "Brenda Halafax",
// 			number: 311541514,
// 		},
// 		recipient: user{
// 			name:   "Sally Sue",
// 			number: 24325,
// 		},
// 	})
// 	test(messageToSend{
// 		message: "you have an event tomorrow",
// 		sender: user{
// 			number: 2352355,
// 		},
// 		recipient: user{
// 			name:   "Suzie Sall",
// 			number: 0,
// 		},
// 	})
// 	test(messageToSend{
// 		message: "you have a party tomorrow",
// 		sender: user{
// 			name:   "Njorn Halafax",
// 			number: 32353125311,
// 		},
// 		recipient: user{
// 			name:   "Sally Sue",
// 			number: 24325,
// 		},
// 	})
// 	test(messageToSend{
// 		message: "you have a birthday tomorrow",
// 		sender: user{
// 			name:   "Eli Halafax",
// 			number: 0,
// 		},
// 		recipient: user{
// 			name:   "Whitakes Sue",
// 			number: 24325,
// 		},
// 	})
// }

//// 14:38
// Anonymous structs in Go

// An anonymous struct is just like a normal struct, but it is defined without a
// name and therefore cannot be referenced elsewhere in the code.

// // To create an anonymous struct, just instantiate the instance immediately
// // using a second pair of brackets afeter declaring the type:
// myCar := struct {
// 	Make string
// 	Model string
// } {
// 	Make: "tesla",
// 	Model: "model 3"
// }

// // You can even nest anontmous structs as fields within other structs:
// type car struct {
// 	Make string
// 	Model string
// 	Height int
// 	Width int
// 	// Wheel is a field containing an anonymous struct
// 	Wheel struct {
// 		Radius int
// 		Material string
// 	}
// }

// In general, PREFER NAMED STRUCTS.

// 14:42
// Embedded structs

// // Go is not an object oriented language. However, embedded structs provide
// // a kind of data-only inheritance that can be useful at times. Keep in mind,
// // GO doesn't support classes or inheritance in complete sense,
// // embedded structs are just a way to elevate adn share fields between struct
// // definitions
// type car struct {
// 	make  string
// 	model string
// }

// type truck struct {
// 	// "car" is embedded, so the definition of a
// 	// "truck" now additionally contains all
// 	// of the fields of the car struct
// 	car
// 	bedSize int
// }

// // Embedded vs Nested
// // * An embedded struct's fields are accessed at the top-level, unlike nested
// // structs.
// // * Prmoted fields can be accessed like normal fields except that they can't
// // be used in composite literals
// lanesTruck := truck{
// 	bedSize: 10,
// 	car: car{
// 		make:  "toyota",
// 		model: "camry",
// 	},
// }

// fmt.Println(lanesTruck.bedSize)

// // embedded fields promoted to the top-level
// // instead of lanesTruct.car.make
// fmt.Println(lanesTruck.make)
// fmt.Println(lanesTruck.model)

// package main

// import "fmt"

// type sender struct {
// 	rateLimit int
// 	user
// }

// type user struct {
// 	name   string
// 	number int
// }

// func test(s sender) {
// 	fmt.Println("Sender name:", s.name)
// 	fmt.Println("Sender number:", s.number)
// 	fmt.Println("Sender rateLimit:", s.rateLimit)
// 	fmt.Println("====================================")
// }

// func main() {
// 	test(sender{
// 		rateLimit: 1000,
// 		user: user{
// 			name:   "Deborah",
// 			number: 123,
// 		},
// 	})
// 	test(sender{
// 		rateLimit: 5000,
// 		user: user{
// 			name:   "Sarah",
// 			number: 1234,
// 		},
// 	})
// 	test(sender{
// 		rateLimit: 1000,
// 		user: user{
// 			name:   "Sally",
// 			number: 12345,
// 		},
// 	})
// }

//// 14:57
// Struct Methods in Go
// Structs in Go kiiiinda look like classes in OOP languages

// While Go is NOT object oriented, it does support methods which can be
// defined on structs. Methods are jsut functions that have a receiver.
// A receiver is a special parameter that syntactially
// goes before the name of the function

// type rect struct {
// 	width int
// 	height int
// }

// // area has a receiver of (r rect)
// func (r rect) area() int {
// 	return r.width * r.height
// }

// r := rect{
// 	widht: 5,e
// 	height: 10,
// }

// fmt.Println(r.area())
// // prints 50

// // A receiver is just a special kind of function parameter. Receivers are
// // important because they will, as you'll learn in the exercies to come, allow
// // us to define interfaces that our structs (and other types) can implement.

// Assignment

package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (authI authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf(
		"Authorization: Basic %s:%s",
		authI.username,
		authI.password,
	)
}

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("====================================")
}

func main() {
	test(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	test(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	test(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}
