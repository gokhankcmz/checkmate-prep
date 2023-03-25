package main

import (
	"cmDeneme/example"
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	user := example.User{
		Name:        "",
		Surname:     "",
		YearOfBirth: 0,
		Address: example.Address{
			City:    "",
			Country: nil,
		},
		Email:   "",
		WebPage: "",
	}

	errs := user.Validate()
	for _, v := range errs {
		fmt.Println(v.Error())
	}
	fmt.Println("Total execution time: ", time.Since(start))
}
