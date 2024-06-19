package request

type Pagenation struct {
	Page     int64
	PageSize int64
}

type SocpeOrder struct {
	OrderBy string
	Order   string `enums:"asc,desc"`
}
