package aria2

import (
	"os/exec"
	"syscall"
)

func terminateProcess(pid int) {
	handle, err := syscall.OpenProcess(syscall.PROCESS_TERMINATE, true, uint32(pid))
	if err != nil {
		return
	}

	syscall.TerminateProcess(handle, 0)
	syscall.CloseHandle(handle)
}

type Aria2Cmd struct {
	*exec.Cmd
}

func NewCmd(path string, args []string) *Aria2Cmd {
	return &Aria2Cmd{
		exec.Command(path, args...),
	}
}

func (c *Aria2Cmd) Close() error {
	// err := c.Process.Signal(os.Interrupt)
	// if err != nil {
	// 	return err
	// }
	err := c.Process.Kill()
	if err != nil {
		return err
	}
	err = c.Process.Release()
	if err != nil {
		return err
	}

	terminateProcess(c.Process.Pid)

	return nil
}
