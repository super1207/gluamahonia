# LuaMahonia for GopherLua

A native Go implementation of [mahonia](https://github.com/axgle/mahonia) library warp for the [GopherLua](https://github.com/yuin/gopher-lua) VM.

## Using

### Loading Modules

```go
import (
	"github.com/super1207/gluamahonia"
)

// Bring up a GopherLua VM
L := lua.NewState()
defer L.Close()

// Preload LuaMahonia modules
gluamahonia.Preload(L) // or gluamahonia.PreloadWithName(L, "mahonia")
```

### Lua example

```lua
local mahonia = require("mahonia")
local utf8Str = "你好"
local gbkStr = mahonia.convert("utf-8","gbk",utf8Str)
print(gbkStr) -- 浣犲ソ
```

## License

MIT