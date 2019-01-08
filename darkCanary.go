package main

import (
  "fmt"
  "os/exec"
)


func main() {
  cmd := exec.Command("/Applications/Canary.app/Contents/MacOS/Google Chrome Canary", "--force-dark-mode")
  stdoutStderr, err := cmd.CombinedOutput()
  if err != nil {
    panic(err)
  }
  fmt.Printf("%s\n", stdoutStderr)
}
