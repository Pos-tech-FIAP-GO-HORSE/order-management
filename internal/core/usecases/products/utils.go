package products

func normalizePage(page int64) int64 {
	if page <= 0 {
		page = 1
	}

	return page
}

func normalizeLimit(limit int64) int64 {
	if limit <= 0 {
		limit = 20
	}

	if limit > 100 {
		limit = 100
	}

	return limit
}

func calculateOffset(page, limit int64) int64 {
	return (page - 1) * limit
}
