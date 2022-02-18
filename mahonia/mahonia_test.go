package gluamahonia_mahonia_test

import (
	"testing"

	"github.com/super1207/gluamahonia"
	lua "github.com/yuin/gopher-lua"
)

func TestMahonia(t *testing.T) {
	L := lua.NewState()
	defer L.Close()
	gluamahonia.Preload(L)
	script := `
		mahonia = require("mahonia")
		return mahonia.convert("utf-8","gbk","你好")
	`
	want := "浣犲ソ"
	if L.DoString(script) != nil {
		t.Fatal("Err in DoString")
	}
	got := lua.LVAsString(L.Get(1))
	if got != want {
		t.Fatalf("Want `%q` but got `%q`", want, got)
	}
}
