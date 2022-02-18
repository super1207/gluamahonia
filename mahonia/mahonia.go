package gluamahonia_mahonia

import (
	mahonia "github.com/axgle/mahonia"
	lua "github.com/yuin/gopher-lua"
)

var exports = map[string]lua.LGFunction{
	"convert": convertFn,
}

func convertFn(L *lua.LState) int {
	s1 := lua.LVAsString(L.Get(1))
	s2 := lua.LVAsString(L.Get(2))
	s3 := lua.LVAsString(L.Get(3))
	srcDecoder := mahonia.NewDecoder(s1)
	desDecoder := mahonia.NewDecoder(s2)
	resStr := srcDecoder.ConvertString(s3)
	_, resBytes, err := desDecoder.Translate([]byte(resStr), true)
	if err != nil {
		L.Push(lua.LString(""))
		return 1
	}
	L.Push(lua.LString(resBytes))
	return 1
}

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.Push(mod)

	L.SetField(mod, "_DEBUG", lua.LBool(false))
	L.SetField(mod, "_VERSION", lua.LString("0.0.0"))

	// consts
	L.SetField(mod, "RAW_DATA", lua.LNumber(1))
	L.SetField(mod, "ZERO_PADDING", lua.LNumber(2))

	return 1
}
