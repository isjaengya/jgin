package common

import "github.com/gin-gonic/gin/binding"

func init() {
	binding.Validator = new(defaultValidator)
}

var modelTableNameMap = map[string]string{
	"Category": "categories",
}
