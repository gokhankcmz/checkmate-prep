package main

import (
	"cmDeneme/types"
	"fmt"
)

func main() {

	user := types.User{
		Name:    "",
		Address: types.Address{},
	}
	errs := user.Validate()
	for _, v := range errs {
		fmt.Println(v.Error())
	}
}
