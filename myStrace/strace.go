package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	var err error
	var regs syscall.PtraceRegs
	var ss syscallCounter
	ss = ss.init()

	fmt.Println("Run: ", os.Args[1:])

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ptrace: true,
	}

	cmd.Start()
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Wait err %v \n", err)
	}

	pid := cmd.Process.Pid
	exit := true

	for {
		if exit {
			err = syscall.PtraceGetRegs(pid, &regs)
			if err != nil {
				break
			}
			//fmt.Printf("%#v \n",regs)
			name := ss.getName(regs.Orig_rax)
			fmt.Printf("name: %s, id: %d \n", name, regs.Orig_rax)
			ss.inc(regs.Orig_rax)
		}

		/**
		http://www.linuxjournal.com/article/6100?page=0,1
		Here we are tracing the write system calls, and ls makes three write system calls. The call to ptrace, with a first argument of PTRACE_SYSCALL, makes the kernel stop the child process whenever a system call entry or exit is made. It's equivalent to doing a PTRACE_CONT and stopping at the next system call entry/exit.
		*/
		err = syscall.PtraceSyscall(pid, 0)
		if err != nil {
			panic(err)
		}

		// http://www.linuxjournal.com/article/6100?page=0,1
		//The status variable in the wait call is used to check whether the child has exited. This is the typical way to check whether the child has been stopped by ptrace or was able to exit.
		_, err = syscall.Wait4(pid, nil, 0, nil)
		if err != nil {
			panic(err)
		}

		exit = !exit
	}

	ss.print()
}
