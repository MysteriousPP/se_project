package v1

import (
	"se_project/internal/model/input/healthin"

	"github.com/gogf/gf/v2/frame/g"
)

type UploadBpReq struct {
	g.Meta `path:"/uploadBp" method:"post" tags:"health" summary:"上传bp接口"`
	healthin.UploadBpInp
}

type UploadBpRes struct {
	*healthin.UploadBpModel
}
