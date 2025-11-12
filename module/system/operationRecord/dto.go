package operationRecord

import "bosh-admin/module/basic"

type GetOperationRecordListReq struct {
	basic.Pagination
	Username  string `json:"username" form:"username"`
	Method    string `json:"method" form:"method"`
	Path      string `json:"path" form:"path"`
	Status    int    `json:"status" form:"status"`
	RequestIP string `json:"requestIP" form:"requestIP"`
	StartTime string `json:"startTime" form:"startTime"`
	EndTime   string `json:"endTime" form:"endTime"`
}

type OperationRecordListItem struct {
	Id             uint   `json:"id"`
	CreatedAt      string `json:"createdAt"`
	Uid            uint   `json:"uid"`
	Username       string `json:"username"`
	Method         string `json:"method"`
	Path           string `json:"path"`
	Status         int    `json:"status"`
	Latency        int64  `json:"latency"`
	RequestIP      string `json:"requestIP"`
	RequestRegion  string `json:"requestRegion"`
	RequestOS      string `json:"requestOS"`
	RequestBrowser string `json:"requestBrowser"`
}
