package easyapi

import (
	"testing"
)

const (
	errCode    = "1"
	errText    = "Fatal"
	errDetails = "Testing"
)

func errorToString(err error) string {
	return err.Error()
}

func TestApiError(t *testing.T) {
	a_err := NewError(errCode, errText, errDetails)
	if errCode != errorToString(a_err) {
		t.Error("Неверный код ошибки или не совместимость с типом error")
	}
	if a_err.Status != errStatus {
		t.Error("Неверный статус")
	}
	if a_err.Text != errText {
		t.Error("Неверный текст")
	}
	if a_err.Details != errDetails {
		t.Error("Неверное описание")
	}
}
