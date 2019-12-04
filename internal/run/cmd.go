package run

import (
	"io"
	"os"
	"os/exec"
	"syscall"
)

// The cmd interface provides a generic form of exec.Cmd so that it can be mocked out in tests.
type cmd interface {
	// methods that exist on exec.Cmd
	Output() ([]byte, error)
	Run() error
	CombinedOutput() ([]byte, error)
	Start() error
	StderrPipe() (io.ReadCloser, error)
	StdinPipe() (io.WriteCloser, error)
	StdoutPipe() (io.ReadCloser, error)
	String() string
	Wait() error

	// setters for struct fields set by callers of exec.Cmd
	Path(string)
	Args([]string)
	Env([]string)
	Dir(string)
	Stdin(io.Reader)
	Stdout(io.Writer)
	Stderr(io.Writer)
	ExtraFiles([]*os.File)
	SysProcAttr(*syscall.SysProcAttr)

	// getters for struct fields generated by exec.Cmd
	Process() *os.Process
	ProcessState() *os.ProcessState
}

// execCmd wraps exec.Cmd along with setters for exec.Cmd's struct fields, implementing the cmd interface.
type execCmd struct {
	*exec.Cmd
}

func (c *execCmd) Path(p string)                      { c.Cmd.Path = p }
func (c *execCmd) Args(a []string)                    { c.Cmd.Args = a }
func (c *execCmd) Env(e []string)                     { c.Cmd.Env = e }
func (c *execCmd) Dir(d string)                       { c.Cmd.Dir = d }
func (c *execCmd) Stdin(s io.Reader)                  { c.Cmd.Stdin = s }
func (c *execCmd) Stdout(s io.Writer)                 { c.Cmd.Stdout = s }
func (c *execCmd) Stderr(s io.Writer)                 { c.Cmd.Stderr = s }
func (c *execCmd) ExtraFiles(f []*os.File)            { c.Cmd.ExtraFiles = f }
func (c *execCmd) SysProcAttr(s *syscall.SysProcAttr) { c.Cmd.SysProcAttr = s }

func (c *execCmd) Process() *os.Process           { return c.Cmd.Process }
func (c *execCmd) ProcessState() *os.ProcessState { return c.Cmd.ProcessState }
