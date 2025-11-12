package loginRecord

import "bosh-admin/module/basic"

type GetLoginRecordListReq struct {
	basic.Pagination
	Username  string `json:"username" form:"username"`
	StartTime string `json:"startTime" form:"startTime"`
	EndTime   string `json:"endTime" form:"endTime"`
	Status    *int   `json:"status" form:"status"`
}
