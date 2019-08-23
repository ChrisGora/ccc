// NOTE: You should not have to modify any functions or methods in this file (other than to add debugging prints).
// Feel free to add other methods and functions as you see fit.
package main

import (
	"container/list"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// bank is made of a number of different accounts.
type bank struct {
	accounts               []*account
	moneyTransferred       int
	transactionsInProgress *list.List // list of transactions
	mutex                  sync.Mutex
	lockedBy               string
	gen                    *dotGenerator
}

// account has an account number, the balance and a mutex lock.
type account struct {
	name     string
	balance  int
	mutex    sync.Mutex
	locked   bool
	lockedBy string
}

// transaction describes an amount to be transferred between two account numbers.
type transaction struct {
	from, to, amount int
	handledBy        string
}

// randomInt generates a random integer between 0 and n.
func randomInt(n int) int {
	x := rand.Intn(n)
	return x
}

// sum return the sum of all account balances in the bank.
func (bank *bank) sum() (sum int) {
	for _, a := range bank.accounts {
		sum += a.balance
	}
	return
}

// lock locks the bank mutex.
func (bank *bank) lock(lockedBy string) {
	bank.mutex.Lock()
	bank.lockedBy = lockedBy
	//if *debug {
	//	fmt.Println("Bank locked by", lockedBy)
	//}
}

// unlock unlocks the bank mutex.
func (bank *bank) unlock() {
	if *debug {
		bank.gen.export(bank)
	}
	//unlockedBy := bank.lockedBy
	bank.lockedBy = ""
	bank.mutex.Unlock()
	//if *debug {
	//	fmt.Println("Bank unlocked by", unlockedBy)
	//}
}

// lockAccount locks the account mutex.
// Call this function before every transaction.
func (bank *bank) lockAccount(accountNumber int, lockedBy string) {
	account := bank.accounts[accountNumber]
	account.mutex.Lock()
	bank.lock(lockedBy)
	account.locked = true
	account.lockedBy = lockedBy
	bank.unlock()
}

// unlockAccount unlocks the account mutex.
// Call this function after every transaction.
func (bank *bank) unlockAccount(accountNumber int, lockedBy string) {
	account := bank.accounts[accountNumber]
	account.mutex.Unlock()
	bank.lock(lockedBy)
	account.locked = false
	account.lockedBy = ""
	bank.unlock()
}

// addInProgress adds a new transaction to the licked list of all active transactions.
// Returns an element of the list 'e' that must be removed once the transaction has finished.
func (bank *bank) addInProgress(transaction transaction, executorId int) *list.Element {
	bank.lock("Executor " + strconv.Itoa(executorId))
	transaction.handledBy = strconv.Itoa(executorId)
	e := bank.transactionsInProgress.PushBack(transaction)
	bank.unlock()
	return e
}

// removeCompleted removes a completed transaction from the licked list of all active transactions.
func (bank *bank) removeCompleted(element *list.Element, executorId int) {
	bank.lock("Executor " + strconv.Itoa(executorId))
	bank.transactionsInProgress.Remove(element)
	bank.unlock()
}

// getTransaction generates a new transaction to be handled by the executor.
// In reality this would come from the bank customers and staff.
// YOU ARE NOT ALLOWED TO MODIFY THIS FUNCTION
func (bank *bank) getTransaction() transaction {
	from := randomInt(len(bank.accounts))
	to := from
	for to == from {
		to = randomInt(len(bank.accounts))
	}
	amount := randomInt(100)
	if to == from {
		panic("incorrect transaction")
	}
	return transaction{from: from, to: to, amount: amount}
}

// execute applies a transaction to the bank system.
// It subtracts the balance from one account and adds it to another, as described in the transaction struct.
// YOU ARE NOT ALLOWED TO MODIFY THIS FUNCTION
func (bank *bank) execute(t transaction, executorId int) {
	bank.lock("Executor " + strconv.Itoa(executorId))
	balanceFrom := bank.accounts[t.from].balance
	balanceTo := bank.accounts[t.to].balance
	bank.unlock()

	duration := time.Duration(50 + randomInt(100))
	time.Sleep(duration * time.Millisecond)

	bank.lock("Executor " + strconv.Itoa(executorId))
	bank.accounts[t.from].balance = balanceFrom - t.amount
	bank.accounts[t.to].balance = balanceTo + t.amount
	bank.moneyTransferred += t.amount
	bank.unlock()
}

func (bank *bank) getAccountName(accountNumber int) string {
	return bank.accounts[accountNumber].name
}
