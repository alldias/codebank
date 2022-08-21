package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type TransactionRepository interface {
	SaveTransaction(transaction Transaction, creditCard CreditCard) error
	GetCreditCard(creditCard CreditCard) (CreditCard, error)
	CreateCreditCard(creditCard CreditCard) error
}

type Transaction struct {
	ID string
	Amount float64
	Status string
	Description string
	Store string 
	CreditCardId string
	CreatedAt time.Time
}

func NewTransaction() *Transaction {
	t := &Transaction{}
	t.ID = uuiid.NewV4().String()
	t.CreatedAt = time.Now()
	return t
}

func (t *Transaction) ProcessAndValidate(createdCard *createdCard) {
	if t.Amount + createdCard.Balance > createdCard.Limit {
		t.Status = "rejected"
	} else {
		t.Status = "approved"
		createdCard.Balance = createdCard.Balance + t.Amount
	}
}