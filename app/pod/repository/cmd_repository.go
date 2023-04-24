package repository

import (
	"os/exec"
	"sync"
)

type CmdMap struct {
	sync.Mutex
	Map map[string]*exec.Cmd
}

func NewCmdMap() *CmdMap {
	return &CmdMap{
		Map: make(map[string]*exec.Cmd),
	}
}

var cmdMapInstance *CmdMap
var cmdMapOnce sync.Once

func (r CmdMap) GetCmdMapInstance() *CmdMap {
	cmdMapOnce.Do(func() {
		cmdMapInstance = &CmdMap{
			Map: make(map[string]*exec.Cmd),
		}
	})
	return cmdMapInstance
}
