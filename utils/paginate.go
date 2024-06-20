package utils

// Paginate is a utility function to paginate data based on page number and limit.

func Paginate[T any](data []T, page int, limit int) []T {
	start := (page - 1) * limit
	end := page * limit

	// Checking bounds to avoid out-of-range errors
	if start >= len(data) {
		return []T{}
	}
	if end > len(data) {
		end = len(data)
	}

	return data[start:end]
}

// func Paginate[T any](data []T, page int, limit int) []T {
// 	// Validate page and limit parameters
// 	if page <= 0 {
// 		page = 1 // Set default page to 1 if page is non-positive
// 	}
// 	if limit <= 0 {
// 		limit = 10 // Set default limit to 10 if limit is non-positive
// 	}

// 	start := (page - 1) * limit
// 	end := page * limit

// 	// Checking bounds to avoid out-of-range errors
// 	if start >= len(data) {
// 		return []T{}
// 	}
// 	if end > len(data) {
// 		end = len(data)
// 	}

// 	return data[start:end]
// }
