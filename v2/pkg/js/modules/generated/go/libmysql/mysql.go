package mysql

import (
	lib_mysql "github.com/projectdiscovery/nuclei/v2/pkg/js/libs/mysql"

	"github.com/dop251/goja"
	"github.com/projectdiscovery/nuclei/v2/pkg/js/gojs"
)

var (
	module = gojs.NewGojaModule("nuclei/mysql")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions

			// Var and consts

			// Types (value type)
			"Client": func() lib_mysql.Client { return lib_mysql.Client{} },

			// Types (pointer type)
			"NewClient": func() *lib_mysql.Client { return &lib_mysql.Client{} },
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}