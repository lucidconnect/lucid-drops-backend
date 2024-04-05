package structure

type ErrorCode string

const (
	LucidInternalError  ErrorCode = "LUCID_INTERNAL_ERROR"
	InvalidRequestError ErrorCode = "LUCID_REQUEST_INVALID"
)
