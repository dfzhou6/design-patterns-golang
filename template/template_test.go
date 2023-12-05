package template

import "testing"

func TestTemplate(t *testing.T) {
	handler := NewFileHandler(NewXmlFile())
	handler.Handle()
	handler.SetProcessor(NewCsvFile())
	handler.Handle()
}
