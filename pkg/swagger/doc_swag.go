package swagger

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

var SwagHandler gin.HandlerFunc

func init() {

	SwagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler)
}
