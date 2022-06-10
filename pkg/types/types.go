package types

type Money int64

type Category string

type Status string

const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

type Payment struct {
	ID int64
	AccountID int64
	Amount Money
	Category Category
	Status Status
}

type Phone string

type Account struct {
	ID int64
	Phone Phone
	Balance Money
}