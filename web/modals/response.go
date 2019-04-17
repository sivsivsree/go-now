package modals

type ResErr struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func (err *ResErr) ErrorMessage(errorMsg string) *ResErr {
	if errorMsg == "" {
		errorMsg = "Not Found"
	}

	err.Error = errorMsg
	return err
}
