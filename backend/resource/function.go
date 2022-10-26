package resource

import "fmt"

func PrintIfErrorExist(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func createUser()
