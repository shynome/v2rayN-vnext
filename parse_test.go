package vnext_test

import (
	"testing"

	. "github.com/shynome/v2rayN-vnext"
)

const (
	vnext_vmess = "vmess://ew0KICAidiI6ICIyIiwNCiAgInBzIjogIjEyNy4wLjAuMSIsDQogICJhZGQiOiAiMTI3LjAuMC4xIiwNCiAgInBvcnQiOiAiNTAwMiIsDQogICJpZCI6ICJlZjc4ZjFkNi01ZmNhLTRjODQtYWNmMi05MGMxZDJiMzhjM2IiLA0KICAiYWlkIjogIjY0IiwNCiAgIm5ldCI6ICIiLA0KICAidHlwZSI6ICIiLA0KICAiaG9zdCI6ICIiLA0KICAicGF0aCI6ICIiLA0KICAidGxzIjogIiINCn0="
	vnext_ss    = "ss://YWVzLTI1Ni1jZmI6MTIzNDU2Nzg5QDEyNy4wLjAuMTo0MDQ=#test+ss+notwork"
)

func TestParseSS(t *testing.T) {
	// 暂时不打算支持
	if true {
		return
	}
	config, err := New(vnext_ss)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(config)
}

func TestParseVMess(t *testing.T) {
	vnext, err := New(vnext_vmess)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(vnext)
}
