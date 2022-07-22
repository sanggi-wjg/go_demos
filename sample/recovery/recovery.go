package recovery

import (
	"errors"
	"fmt"
)

func panicRecover() {
	defer func() {
		r := recover()
		fmt.Println("Recover panic: ", r)
	}()
	panic(errors.New("something is wrong..."))
	// unreachable
	// fmt.Println("panicRecover() is done.")
}

func RecoveryMain() {
	fmt.Println("START!")
	panicRecover()
	fmt.Print("RecoveryMain() is done.")
}

/*
START!
Recover panic:  something is wrong...
RecoveryMain() is done.
*/
