package command

import "fmt"

type Command interface {
	Exec()
}

type SqlCommand struct {
	sql string
}

func NewSqlCommand(sql string) *SqlCommand {
	return &SqlCommand{sql: sql}
}

func (impl *SqlCommand) Exec() {
	fmt.Printf("sql command exec, %v\n", impl.sql)
}

type Request struct {
	Body   string
	Method string
}

type RequestHandlerCommand struct {
	req *Request
}

func NewRequestHandlerCommand(req *Request) *RequestHandlerCommand {
	return &RequestHandlerCommand{req: req}
}

func (impl *RequestHandlerCommand) Exec() {
	fmt.Printf("request handler command exec, request is %+v\n", *impl.req)
}

type CommandRunner struct {
	commands []Command
}

func NewCommandRunner() *CommandRunner {
	return &CommandRunner{}
}

func (impl *CommandRunner) AddCommand(command Command) {
	impl.commands = append(impl.commands, command)
}

func (impl *CommandRunner) RunAllCommand() {
	for _, item := range impl.commands {
		item.Exec()
	}
}
