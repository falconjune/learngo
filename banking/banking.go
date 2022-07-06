package banking

import "errors"

//Account struct
type Account struct {
	owner   string
	balance int
}

//모듈을 불러와서 실행시킬때는 go mod init을 통해 mod를 go.mod 파일을 만든다.
func NewAccount(owner1 string) *Account {
	account := Account{owner: owner1, balance: 0}
	return &account
}

var noMoney = errors.New("Can't withdraw")

//reciever를 통해 method와 비슷한 기능을 하는 함수를 제작함.
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//balance만을 보여주는 함수를 작성함.
func (a Account) Balance() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return noMoney
	}
	a.balance -= amount
	return nil
}

func (a *Account) Changeowner(cowner string) {
	a.owner = cowner
}
