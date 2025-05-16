package models

const limit = 10

func GetDefaultPaginationInputId() *PaginationInputId {
	return &PaginationInputId{
		Limit: limit,
	}
}

func GetDefaultPaginationInputInt() *PaginationInputInt {
	return &PaginationInputInt{
		Limit: limit,
	}
}

func GetDefaultPaginationInputString() *PaginationInputString {
	return &PaginationInputString{
		Limit: limit,
	}
}

func GetDefaultPaginationInputTimestamp() *PaginationInputTimestamp {
	return &PaginationInputTimestamp{
		Limit: limit,
	}
}
