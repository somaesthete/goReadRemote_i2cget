package cmd

import (
	"BBB_remotes/pkg/oscmanager"
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/bitfield/script"
	"os"
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

func hexToParallel(byteStr string) {
	sanitizedByteStr := strings.Trim(byteStr, "0x")

	decodedByte, err := hex.DecodeString(sanitizedByteStr)
	if err != nil {
		handle(err)
	}

	b := decodedByte[0]
	fmt.Printf("%t - %t - %t - %t --- %t - %t - %t - %t\n",
		oscmanager.SetSlotEnable(8, (b&0x80)>>7 == 1),
		oscmanager.SetSlotEnable(7, (b&0x40)>>6 == 1),
		oscmanager.SetSlotEnable(6, (b&0x20)>>5 == 1),
		oscmanager.SetSlotEnable(5, (b&0x10)>>4 == 1),
		oscmanager.SetSlotEnable(4, (b&0x8)>>3 == 1),
		oscmanager.SetSlotEnable(3, (b&0x4)>>2 == 1),
		oscmanager.SetSlotEnable(2, (b&0x2)>>1 == 1),
		oscmanager.SetSlotEnable(1, (b&0x1)>>0 == 1),
	)

}

func execWithBufferedReader(scriptPath string) {
	data, err := script.File(scriptPath).String()
	if err != nil {
		handle(err)
	}

	fmt.Println(data)
	os.Exit(0)

	// if we wanted an explicit thing we could just run here.
	//data := "ssh root@192.168.1.132 'while :; do echo $(i2cget -y 2 0x20 0x12); sleep 0.1; done'"

	pipe := script.Exec(data)

	scanner := bufio.NewScanner(pipe.Reader)

	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
		//hexToParallel(m)
	}
}

func handle(err error) {
	fmt.Printf("oops %s", err)
}
