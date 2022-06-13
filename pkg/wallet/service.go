package wallet

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/bahodurnazarov/bankV2/v2/pkg/types"
	"github.com/google/uuid"
)

var ErrPhoneRegistred = errors.New("Phone already registred")
var ErrAmountMustBePosititve = errors.New("amount mus be greater than 0")
var ErrAccountNotFound = errors.New("account not found")
var ErrNotEnoughBalance = errors.New("Not Enough Balance ")

type Service struct {
	nextAccountID int64
	accounts      []*types.Account
	payments      []*types.Payment
}

type Messenger interface {
	Send(message string) bool
	Receive() (message string, ok bool)
}

type Telegram struct {
}

type Error string

func (e Error) Error() string {
	return string(e)
}

func (t *Telegram) Send(message string) bool {
	return true
}

func (t *Telegram) Receive() (message string, ok bool) {
	return "", true
}

func RegisterAccount(service *Service, phone types.Phone) {
	for _, account := range service.accounts {
		if account.Phone == phone {
			return
		}
	}
	service.nextAccountID++
	service.accounts = append(service.accounts, &types.Account{
		ID:      service.nextAccountID,
		Phone:   phone,
		Balance: 0,
	})
}
func (s *Service) RegisterAccount(phone types.Phone) (*types.Account, error) {
	for _, account := range s.accounts {
		if account.Phone == phone {
			return nil, ErrPhoneRegistred
		}
	}
	s.nextAccountID++
	account := &types.Account{
		ID:      s.nextAccountID,
		Phone:   phone,
		Balance: 0,
	}
	s.accounts = append(s.accounts, account)
	return account, nil
}

func (s *Service) Deposit(accountID int64, amount types.Money) error {
	if amount <= 0 {
		return ErrAmountMustBePosititve
	}

	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account = acc
			break
		}
	}

	if account == nil {
		return ErrAccountNotFound
	}

	account.Balance += amount
	return nil
}

func (s *Service) Pay(accountID int64, amount types.Money, category types.Category) (*types.Payment, error) {
	if amount <= 0 {
		return nil, ErrAmountMustBePosititve
	}

	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account = acc
			break
		}
	}
	if account == nil {
		return nil, ErrAccountNotFound
	}

	if account.Balance < amount {
		return nil, ErrNotEnoughBalance
	}

	account.Balance -= amount
	paymentID, err := strconv.ParseInt(uuid.New().String(), 10, 64)
	if err != nil {
		fmt.Println(err)
		return &types.Payment{}, err
	}


	payment := &types.Payment{
		ID:        paymentID,
		AccountID: accountID,
		Amount:    amount,
		Category:  category,
		Status:   types.StatusInProgress,
	}
	s.payments = append(s.payments, payment)
	return payment, nil
}
