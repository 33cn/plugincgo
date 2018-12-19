package pbftlibbyz

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"testing"
)

func TestPbft(t *testing.T) {
	bi := exec.Command("/bin/sh", "-c", "./test/build-docker.sh")
	err := bi.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Docker build success!")
	// 打印log信息
	cmd := exec.Command("/bin/sh", "-c", "./test/run-docker.sh")
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Print(line)
	}
	cmd.Wait()
}
