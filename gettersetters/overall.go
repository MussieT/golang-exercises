package gettersetters

import (
	"fmt"
	"time"
)

// In go, getters and setters are considered neither idiomatic nor mandotory. - example, time has .. time.Timer - the field directly
// If you must have to use Getter .. use Balance, not GetBalance, for the setter .. setBalance (customer.Balance, customer.SetBalance())
func GettersSetters() {
	fmt.Println("starting timer")
	timer := time.NewTimer(time.Second)
	<-timer.C
	fmt.Println("the heck")
}
// READ ABOUT GETTERS AND SETTERS ..
// CONTROVERSIES ABOUT GETTERS -- SETTERS 

//  Here's the core of the issue: if you have a bunch of private member variables, 
//  but you expose them via getters/setters, then you may as well have made them public in the first place. 
//  In that sense, you've got zero encapsulation. So yeah, it's a bad practice in that sense. 

// Advantage
// For example, imagine you have a class called Person that has a field called age. 
// You won’t want this to be directly changeable from outside the class, as someone could set the age to an invalid value such as -3. 
// func setAge(int value) {
//     if (age < 0)
//         throw new IllegalArgumentException();
//     age = value;
// }
