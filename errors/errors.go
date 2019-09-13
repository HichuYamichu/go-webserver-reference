package errors

// AppError : custom htpp error type
type ApiError struct {
	Err  error
	Msg  string
	Code int
}
