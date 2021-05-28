/*
Copyright 2021 Alexaps
*/

package main

import (
	"fmt"

	rhsm "github.com/alexapps/rhsm/licensecontroller"
)

var (
	UUID     = ""
	user     = ""
	password = ""
)

func main() {

	rhsmClient := rhsm.NewRHSMClient()
	rhsmClient.SetCredentials(user, password)
	entitlement, err := rhsmClient.GetEntitlement(UUID)
	if err != nil {
		fmt.Println("Entitlement getting error", err)
	}
	if err := rhsmClient.Delete(UUID, entitlement); err != nil {
		fmt.Println("Delete subscription error", err)
	}

}
