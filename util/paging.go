package util

type Paging struct {
	Page       int `json:"page"`
	Size       int `json:"size"`
	TotalPages int `json:"totalPages"`
	TotalRows  int `json:"totalRows"`
}

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PagingResponse struct {
	Status Status `json:"status"`
	Data   []any  `json:"data"`
	Paging Paging `json:"paging"`
}

type SingleResponse struct {
	Status Status `json:"status"`
	Data   any    `json:"data"`
}
