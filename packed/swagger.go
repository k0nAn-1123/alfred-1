package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GDgqbgcxIAERBg4GYrLE9PTU4v0obReVnF+XmgIKwNjJ9u5hFXPggICAgK09Px9/P11gzYFBRhdMg641NiR9mzpU+1ITYcrDa6TPn/2Eprs7dC0NeFi1IpLTUcm9SR1FOgqMDD8/x/gzc5h1jlhoQMDA0MIAwMDzCkMDAJoTmFHOAVs+wS2cwkg3chqArwZmUSYEV5BNhnkFRjY1ggi8XoMYRR2p0CAAMN/x26EUUgOY2UDyTMxMDF0MjAwnASrBgQAAP//D3LwV2gBAAA="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
