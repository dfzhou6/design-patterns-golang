package command

import "testing"

func TestCommand(t *testing.T) {
	selectSqlCmd := NewSqlCommand("select * from tab")
	updateSqlCmd := NewSqlCommand("update tab set name = 'felix'")
	postRequestCmd := NewRequestHandlerCommand(&Request{Body: `{"name":"felix","age":20}`, Method: "POST"})
	getRequestCmd := NewRequestHandlerCommand(&Request{Body: ``, Method: "GET"})
	cmdRunner := NewCommandRunner()
	cmdRunner.AddCommand(selectSqlCmd)
	cmdRunner.AddCommand(updateSqlCmd)
	cmdRunner.AddCommand(postRequestCmd)
	cmdRunner.AddCommand(getRequestCmd)
	cmdRunner.RunAllCommand()
}
