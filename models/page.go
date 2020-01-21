package models

type QueryParams struct {
	PageNo   int                    `json:"pageNo"`
	PageSize int                    `json:"pageSize"`
	Query    map[string]interface{} `json:"query"`
	SortBy   string                 `json:"sortBy"`
	GroupBy  []string               `json:"groupBy"`
}

type Page struct {
	PageNo     int         `json:"pageNo"`
	PageSize   int         `json:"pageSize"`
	TotalPage  int         `json:"totalPage"`
	TotalCount int         `json:"totalCount"`
	List       interface{} `json:"data"`
}

func (q QueryParams) NewPage(count int, list interface{}) Page {
	tp := count / q.PageSize
	if count%q.PageSize > 0 {
		tp = count/q.PageSize + 1
	}
	return Page{PageNo: q.PageNo, PageSize: q.PageSize, TotalPage: tp, TotalCount: count, List: list}
}
