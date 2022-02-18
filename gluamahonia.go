package gluamahonia

import (
	mahonia "github.com/super1207/gluamahonia/mahonia"
	lua "github.com/yuin/gopher-lua"
)

func Preload(L *lua.LState) {
	PreloadWithName(L, "mahonia")
}

func PreloadWithName(L *lua.LState, name string) {
	L.PreloadModule(name, mahonia.Loader)
}
