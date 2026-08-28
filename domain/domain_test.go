package domain

import (
	"testing"
	"time"
)

func TestPolicy(t *testing.T) {
	p := DefaultPolicy()
	if !p.PrefixAllowed("TRN") || p.SessionFresh(1, 2) == false {
		t.Fatal()
	}
	if (AccessCard{Status: CardActive, ExpiresAt: 2}).Usable(time.Unix(3, 0)) {
		t.Fatal()
	}
}
