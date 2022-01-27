package metadata

import (
	metaReq "github.com/dolica/mds/model/metadata/request"
	"github.com/gin-gonic/gin"
)

type SaleOrgApi struct{}

func (*SaleOrgApi) AddOrg(c *gin.Context) {
	var scopeRequest metaReq.SaleOrgRequest
	c.ShouldBindJSON(scopeRequest)
}

// 查询是否存在，如不存在返回异常响应码
func (*SaleOrgApi) GetOrg(c *gin.Context) {

}
