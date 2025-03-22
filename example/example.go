package example

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Signal"
)

type MyGoNode struct {
	classdb.Extension[MyGoNode, Node.Instance] `gd:"MyGoNode"`
	MySignalEventHandler                       Signal.Void              `gd:"my_signal_event_handler()"`
	MySignalWithParamsEventHandler             Signal.Pair[string, int] `gd:"my_signal_with_params_handler(message, n)"`
	MyProperty                                 string
}

func (m *MyGoNode) PrintNodeName(MyGoNode Node.Instance) {
	Engine.Print(MyGoNode.Name())
}

func (m *MyGoNode) PrintArray(arr []string) {
	for _, element := range arr {
		Engine.Print(element)
	}
}

func (m *MyGoNode) PrintNTimes(msg string, n int) {
	for range n {
		Engine.Print(msg)
	}
}

func (m *MyGoNode) MySignalHandler() {
	Engine.Print("The signal handler was called!")
}

func (m *MyGoNode) MySignalWithParamsHandler(msg string, n int) {
	m.PrintNTimes(msg, n)
}
