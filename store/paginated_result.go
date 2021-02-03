package store

// PaginatedResult represents database records with pagination metadata
type PaginatedResult struct {
	Page       int64       `json:"page"`
	Limit      int64       `json:"limit"`
	TotalCount int64       `json:"total_count"`
	TotalPages int64       `json:"total_pages"`
	Records    interface{} `json:"records"`
}

// update calculates the total number of pages based on record count
func (pr *PaginatedResult) update() *PaginatedResult {
	if pr.TotalPages == 0 && pr.TotalCount > 0 {
		totalPages := pr.TotalCount / pr.Limit

		if totalPages*pr.Limit < pr.TotalCount {
			totalPages++
		}

		pr.TotalPages = totalPages
	}

	return pr
}
