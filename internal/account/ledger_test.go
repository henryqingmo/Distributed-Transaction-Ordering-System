package account

import (
	"testing"
)

func TestDeposit(t *testing.T) {
	l := NewLedger()
	l.Deposit("alice", 100)
	if l.balances["alice"] != 100 {
		t.Errorf("expected 100, got %d", l.balances["alice"])
	}
	// second deposit accumulates
	l.Deposit("alice", 50)
	if l.balances["alice"] != 150 {
		t.Errorf("expected 150, got %d", l.balances["alice"])
	}
}

func TestDepositCreatesAccount(t *testing.T) {
	l := NewLedger()
	l.Deposit("newacct", 10)
	if l.balances["newacct"] != 10 {
		t.Errorf("expected 10, got %d", l.balances["newacct"])
	}
}

func TestTransferSuccess(t *testing.T) {
	l := NewLedger()
	l.Deposit("alice", 100)
	ok := l.Transfer("alice", "bob", 40)
	if !ok {
		t.Fatal("expected transfer to succeed")
	}
	if l.balances["alice"] != 60 {
		t.Errorf("alice: expected 60, got %d", l.balances["alice"])
	}
	if l.balances["bob"] != 40 {
		t.Errorf("bob: expected 40, got %d", l.balances["bob"])
	}
}

func TestTransferCreatesDst(t *testing.T) {
	l := NewLedger()
	l.Deposit("alice", 50)
	ok := l.Transfer("alice", "charlie", 50)
	if !ok {
		t.Fatal("expected transfer to succeed")
	}
	if l.balances["charlie"] != 50 {
		t.Errorf("charlie: expected 50, got %d", l.balances["charlie"])
	}
}

func TestTransferInsufficientFunds(t *testing.T) {
	l := NewLedger()
	l.Deposit("alice", 10)
	ok := l.Transfer("alice", "bob", 20)
	if ok {
		t.Fatal("expected transfer to fail")
	}
	// no change
	if l.balances["alice"] != 10 {
		t.Errorf("alice should be unchanged: got %d", l.balances["alice"])
	}
	if l.balances["bob"] != 0 {
		t.Errorf("bob should be 0: got %d", l.balances["bob"])
	}
}

func TestTransferNonExistentSrc(t *testing.T) {
	l := NewLedger()
	ok := l.Transfer("ghost", "bob", 1)
	if ok {
		t.Fatal("expected transfer from non-existent account to fail")
	}
}

func TestBalancesEmpty(t *testing.T) {
	l := NewLedger()
	if l.Balances() != "BALANCES " {
		t.Errorf("unexpected: %q", l.Balances())
	}
}

func TestBalancesOmitsZero(t *testing.T) {
	l := NewLedger()
	l.Deposit("alice", 10)
	l.Transfer("alice", "bob", 10) // alice goes to 0
	got := l.Balances()
	if got != "BALANCES bob:10" {
		t.Errorf("unexpected: %q", got)
	}
}

func TestBalancesSorted(t *testing.T) {
	l := NewLedger()
	l.Deposit("wqkby", 10)
	l.Deposit("yxpqg", 75)
	l.Transfer("yxpqg", "wqkby", 13)
	got := l.Balances()
	want := "BALANCES wqkby:23 yxpqg:62"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// Spec example: two transfers in order, second one rejected
func TestTransferOrderMatters(t *testing.T) {
	l := NewLedger()
	l.Deposit("wqkby", 23)
	l.Deposit("yxpqg", 62)

	ok1 := l.Transfer("wqkby", "hreqp", 20)
	if !ok1 {
		t.Fatal("first transfer should succeed")
	}
	ok2 := l.Transfer("wqkby", "buyqa", 15)
	if ok2 {
		t.Fatal("second transfer should be rejected (insufficient funds)")
	}
	got := l.Balances()
	want := "BALANCES hreqp:20 wqkby:3 yxpqg:62"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
