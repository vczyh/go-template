package log

import (
	"github.com/gin-gonic/gin"
	"io"
)

func configGinLog(w io.Writer)  {
	gin.DefaultWriter = w
}
