package customer

import (
	"fmt"

	"Tutorial/domain"
)

func Update(mObj domain.Customer) (bool, error) {

	fmt.Println("Update Method Invoked!")

	fmt.Println("Customer Record Updated is \t:\t", mObj)

	return true, nil
}
