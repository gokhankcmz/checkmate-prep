package main

import (
	"cmDeneme/checkmate"
	"cmDeneme/checkmate/settings"
	"cmDeneme/types"
	"fmt"
)

func main() {

	checkmate.Init(&settings.Settings{StopAtFirstError: true})

	age := 0
	user := types.User{
		Name:    "this text",
		Address: types.Address{},
		Age:     &age,
	}

	errs := user.Validate()
	for _, v := range errs {
		fmt.Println(v.Error())
	}
}
