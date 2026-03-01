package account

import (
	"fmt"
	"sort"
	"strings"
)



type Ledger struct {
	balances map[string]int
}

func NewLedger() *Ledger {
	return &Ledger{balances: make(map[string]int)}
}

// Deposit adds amount to account. Always succeeds; creates the account if needed.
func (l *Ledger) Deposit(account string, amount int) {
	l.balances[account] += amount
}

// Transfer moves amount from src to dst. Returns false (and makes no change) if
// src does not exist or has insufficient funds. Creates dst if needed.
func (l *Ledger) Transfer(src, dst string, amount int) bool {
	if l.balances[src] < amount {
		return false
	}
	l.balances[src] -= amount
	l.balances[dst] += amount
	return true
}

// Balances returns the BALANCES line as required by the spec:

//	BALANCES a:10 b:20 c:5
// Accounts with zero balance are omitted. Accounts are sorted alphabetically.
func (l *Ledger) Balances() string {
	accounts := make([]string, 0, len(l.balances))
	for a, bal := range l.balances {
		if bal > 0 {
			accounts = append(accounts, a)
		}
	}
	sort.Strings(accounts)

	parts := make([]string, len(accounts))
	for i, a := range accounts {
		parts[i] = fmt.Sprintf("%s:%d", a, l.balances[a])
	}

	return "BALANCES " + strings.Join(parts, " ")
}
