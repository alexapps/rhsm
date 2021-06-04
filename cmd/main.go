/*
Copyright 2021 Alexaps
*/

package main

import (
	"fmt"

	rhsm "github.com/alexapps/rhsm/licensecontroller"
)

var (
	UUID     = "12a9feec-22bf-4de7-b47d-409c3d83f6bb"
	user     = "kaas-ci-rhel"
	password = "OpBauD9ShLxIZN1ccoYR"
)

func main() {

	rhsmClient := rhsm.NewRHSMClient()
	rhsmClient.SetCredentials(user, password)
	entitlement, err := rhsmClient.GetEntitlement(UUID)
	if err != nil {
		fmt.Println("Entitlement getting error", err)
	}
	fmt.Println("Entitlement ", entitlement)

	// Delete subscription
	if err := rhsmClient.DeleteSubscription(UUID, entitlement); err != nil {
		fmt.Println("Delete subscription error", err)
	}
	fmt.Println("Delete subscription done ")

	if err := rhsmClient.DeleteConsumer(UUID); err != nil {
		fmt.Println("Delete consumer error", err)
	}
	fmt.Println("Delete consumer done ")

	if err := rhsmClient.GetConsumer("12a9feec-22bf-4de7-b47d-409c3d83f6bb"); err != nil {
		fmt.Println("Get consumer error", err)
	}
	fmt.Println("Get consumer done ")
}
