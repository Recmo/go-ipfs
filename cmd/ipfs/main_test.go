package main

import (
	"testing"

	"gx/ipfs/QmUb9C21o1apw6UGHUnCK4a2eCoPYHkaRZwXWTYzxcztRU/go-ipfs-cmds/cmdsutil"
)

func TestIsCientErr(t *testing.T) {
	t.Log("Catch both pointers and values")
	if !isClientError(cmdsutil.Error{Code: cmdsutil.ErrClient}) {
		t.Errorf("misidentified value")
	}
	if !isClientError(&cmdsutil.Error{Code: cmdsutil.ErrClient}) {
		t.Errorf("misidentified pointer")
	}
}
