package main

import (
	"fmt"

	"github.com/bahodurnazarov/bankV2/v2/pkg/wallet"
)

func main() {
	svc := &wallet.Service{}
	account, err := svc.RegisterAccount("+99202020202")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.Deposit(account.ID, 10)
	if err != nil {
		switch err {
		case wallet.ErrAmountMustBePosititve:
			fmt.Println(" Сумма должна быть положительной ")
		case wallet.ErrAccountNotFound:
			fmt.Println(" Аккаунт пользователя не найден")
		}
		return
	}

	fmt.Println(account.Balance)
}