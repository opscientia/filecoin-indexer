package store

// Pagination represents the pagination parameters
type Pagination struct {
	Page  int64 `form:"page"`
	Limit int64 `form:"limit"`
}

const (
	_paginationDefaultLimit = 100
	_paginationMaxLimit     = 1000
)

// Validate sanitizes the pagination parameters
func (p *Pagination) Validate() error {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = _paginationDefaultLimit
	}
	if p.Limit >= _paginationMaxLimit {
		p.Limit = _paginationMaxLimit
	}
	return nil
}

func (p Pagination) offset() int {
	return int((p.Page - 1) * p.Limit)
}

func (p Pagination) limit() int {
	return int(p.Limit)
}
