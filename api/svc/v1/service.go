package v1

import (
	"se_project/internal/model/input/svcin"

	"github.com/gogf/gf/v2/frame/g"
)

type ServiceOrderCreateReq struct {
	g.Meta `path:"/service" method:"post" tags:"service" summary:"创建服务订单"`
	svcin.ServiceOrderCreateInp
}

type ServiceOrderCreateRes struct {
	*svcin.ServiceOrderCreateModel
}
