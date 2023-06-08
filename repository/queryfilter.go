package repository

type QueryFilter struct {
	filter string
}

func NewQueryFilter() QueryFilter {
	return QueryFilter{}
}

func (q QueryFilter) AddFilter(filter string, value string) QueryFilter {
	if q.filter == "" {
		q.filter = filter + "=" + value
		return q
	}
	q.filter = q.filter + " AND " + filter + "=" + value
	return q
}

func (q QueryFilter) GetFilter() string {
	return q.filter
}
