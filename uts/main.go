package main
import (
    "os/exec"
    "syscall"
    "os"
    "log"
    
)
func main() {//
    cmd := exec.Command("sh")  //可以理解为进程的名称
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags: syscall.CLONE_NEWUTS,
    }
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err !=nil{
        log.Fatal(err)
    }
}