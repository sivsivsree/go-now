<<<<<<< HEAD:03_creating_containers/main.go
package main
=======
// +build

package lib
>>>>>>> 9377ecb1d8477c0a0057613d4b673acbf6933fa8:lib/createContainer.go

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Create() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("what...")

	}
}
func run() {
	fmt.Printf("Running.. %v as PID %d\n", os.Args[2:], os.Getpid())
	cmd := exec.Command("proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{}

	must(cmd.Run())
}

func child() {
	fmt.Printf("Running.. %v as PID %d\n", os.Args[2:], os.Getpid())
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{}

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}

}
