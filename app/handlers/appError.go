package handlers

// AppError : custom htpp error type
type AppError struct {
	Err  error
	Msg  string
	Code int
}
