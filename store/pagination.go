package store

const (
	paginationDefaultLimit = 100
	paginationMaxLimit     = 1000
)

// Pagination represents the pagination parameters
type Pagination struct {
	Page  int64 `form:"page"`
	Limit int64 `form:"limit"`
}

// Validate sanitizes the pagination parameters
func (p *Pagination) Validate() error {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = paginationDefaultLimit
	}
	if p.Limit >= paginationMaxLimit {
		p.Limit = paginationMaxLimit
	}
	return nil
}

func (p Pagination) offset() int {
	return int((p.Page - 1) * p.Limit)
}

func (p Pagination) limit() int {
	return int(p.Limit)
}
