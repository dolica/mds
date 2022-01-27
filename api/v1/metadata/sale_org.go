package metadata

import (
	metadata "github.com/dolica/mds/model/metadata/request"
	"github.com/gin-gonic/gin"
)

func add(c *gin.Context) {
	var scopeRequest metadata.ScopeRequest
	c.ShouldBindJSON(scopeRequest)

}
