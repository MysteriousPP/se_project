package healthin

import "github.com/gogf/gf/v2/os/gtime"

type OverviewInp struct {
	ElderId int `p:"elderId"`
}

type AbpHistory struct {
	Time  *gtime.Time `json:"time"`
	Value string      `json:"value"`
}
type OverviewModel struct {
	ElderId   int        `json:"elderId"`
	LastestBp string     `json:"lastestBp"`
	BpHistory AbpHistory `json:"bpHistory"`
}
