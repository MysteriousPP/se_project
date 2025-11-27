package v1

import (
	"se_project/internal/model/input/healthin"

	"github.com/gogf/gf/v2/frame/g"
)

type SOSCreateReq struct {
	g.Meta `path:"/sos" method:"post" tags:"health" summary:"创建呼救"`
	healthin.SOSCreateInp
}

type SOSCreateRes struct {
	*healthin.SOSCreateModel
}
