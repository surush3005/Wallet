package main

import (
	"fmt"

	"github.com/surush3005/Wallet/pkg/wallet"
)

func main() {

	svc := &wallet.Service{}
	/*
		wallet.RegisterAccount(svc, "+992918176779")
	*/

	//svc.RegisterAccount("+992918176779")

	account, err := svc.RegisterAccount("+992918176779")
	if err != nil {
		fmt.Println(err)
		fmt.Println(account)	
		return
	}

}
