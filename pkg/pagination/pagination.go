package pagination

type Pagination struct {
	Offset     int
	Limit      int
	Query      string
	Dir        string
	Sort       string
	TotalPages int64
	TotalRows  int64
}

func New(page, limit int, query, dir, sort string) *Pagination {
	return &Pagination{
		Offset: (page - 1) * limit,
		Limit:  limit,
		Query:  query,
		Dir:    dir,
		Sort:   sort,
	}
}

func (p *Pagination) CalculateTotalPage(limit int64, totalRows int64) int64 {
	var totalPage int64 = 1
	if limit > 0 {
		rowMod := totalRows % limit
		if rowMod != 0 {
			rowMod = 1
		}
		totalPage = (totalRows / limit) + rowMod
	}

	return totalPage
}
