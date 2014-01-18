package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/proxypoke/i3ipc"
)

var (
	ipcsocket *i3ipc.IPCSocket
)

type Workspace struct {
	Name    string
	Focused bool
}

func main() {
	events, err := i3ipc.Subscribe(i3ipc.I3WorkspaceEvent)
	if err != nil {
		return
	}

	ipcsocket, err = i3ipc.GetIPCSocket()
	if err != nil {
		return
	}
	printWorkspaces()

	for {
		_ = <-events
		printWorkspaces()
	}
}

func printWorkspaces() {
	resp, err := ipcsocket.Raw(i3ipc.I3GetWorkspaces, "")
	if err != nil {
		return
	}

	var wss []*Workspace
	err = json.Unmarshal(resp, &wss)
	if err != nil {
		return
	}

	var out bytes.Buffer
	out.WriteString("\\c")
	for _, ws := range wss {
		if ws.Focused {
			fmt.Fprintf(&out, "\\u0\\fr  %s  ", ws.Name)
		} else {
			fmt.Fprintf(&out, "\\ur\\fr  %s  ", ws.Name)
		}
	}
	out.WriteString("")

	fmt.Println(string(out.Bytes()))
}
