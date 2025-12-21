package v1

import (
	"se_project/internal/model/input/healthin"

	"github.com/gogf/gf/v2/frame/g"
)

type OverviewReq struct {
	g.Meta `path:"overview" method:"get"`
	healthin.OverviewInp
}

type OverviewRes struct {
	g.Meta `path:"overview" method:"get"`
	*healthin.OverviewModel
}
