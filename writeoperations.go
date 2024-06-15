package customer

import (
	"fmt"

	"Tutorial/domain"
)

func Insert(mObj domain.Customer) (bool, error) {

	fmt.Println("Insert Method Invoked!")

	fmt.Println("Customer Record Inserted is \t:\t", mObj)

	return true, nil
}
