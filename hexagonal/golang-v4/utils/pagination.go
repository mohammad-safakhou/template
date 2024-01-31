package utils

const (
	// DefaultLimit defines the default number of items per page for APIs
	DefaultLimit int = 25

	// DefaultOffset defines the default offset for API responses
	DefaultOffset int = 0
)

func OffsetFromPage(page int, limit int) (offset int) {
	offset = DefaultOffset

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = DefaultLimit
	}

	return (page * limit) - limit
}
