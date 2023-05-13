package main

import (
	"bufio"
	"fmt"
	"github.com/bitfield/script"
	"strings"
)

func main() {
	// depends on gpg priv/pub key
	//cmd := "ssh root@192.168.1.148 'PIN=51; echo 0 > /sys/class/gpio/gpio$PIN/value && sleep 1 && echo 1 > /sys/class/gpio/gpio$PIN/value'"

	//toggleScriptPath := "scripts/toggleGpio.sh"
	//execWithBufferedReader(toggleScriptPath)

	i2cReadPath := "scripts/i2c/scp_i2cGetOnRemote.sh"
	execWithBufferedReader(i2cReadPath)

}

func execWithBufferedReader(scriptPath string) {
	data, err := script.File(scriptPath).String()
	if err != nil {
		handle(err)
	}

	fmt.Println(data)
	//os.Exit(0)

	// if we wanted an explicit thing we could just run here.
	//data := "ssh root@192.168.1.132 'while :; do echo $(i2cget -y 2 0x20 0x12); sleep 0.1; done'"

	for _, cmd := range strings.Split(data, "\n") {
		pipe := script.Exec(cmd)
		scanner := bufio.NewScanner(pipe.Reader)

		for scanner.Scan() {
			m := scanner.Text()
			fmt.Println(m)
			//hexToParallel(m)
		}
	}

}

func handle(err error) {
	fmt.Printf("oops %s", err)
}
