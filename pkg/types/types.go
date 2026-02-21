package types

type Money int64

type PaymentCategory string

type PaymentStatus string

const (
	PaymentStatusOK         PaymentStatus = "OK"
	PaymentStatusFail       PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPRORGRESS"
)

type Payment struct {
	ID        string
	Amount    Money
	AcoountID int64
	Category  PaymentCategory
	Status    PaymentStatus
}

type Phone string

type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}
