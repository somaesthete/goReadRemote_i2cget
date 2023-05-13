package hexToParallel

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

func handle(err error) {
	fmt.Printf("oops %s", err)
}

func HexToParallelSanitizeWord(inputWordRaw string) (wordBool [][]bool) {

	var sanitizedWords string
	targetMatch := regexp.MustCompile(`^0x[0-9a-fA-F]{2} 0x[0-9a-fA-F]{2} 0x[0-9a-fA-F]{2} 0x[0-9a-fA-F]{2}`)

	if !targetMatch.MatchString(inputWordRaw) {
		handle(fmt.Errorf("inputWordRaw %s does not match targetMatch %s", inputWordRaw, targetMatch))
		return nil
	} else {
		sanitizedWords = inputWordRaw
	}

	// now that we have sanitized input
	words := strings.Split(sanitizedWords, " ")

	if len(words) != 4 {
		handle(fmt.Errorf("should be 4 words, got %d", len(words)))
		return nil
	} else {
		boolsArr := [][]bool{HexToParallelOneByte(words[0]),
			HexToParallelOneByte(words[1]),
			HexToParallelOneByte(words[2]),
			HexToParallelOneByte(words[3]),
		}
		return boolsArr
	}
}

func HexToParallelOneByte(byteStr string) (byteBool []bool) {
	sanitizedByteStr := strings.TrimPrefix(byteStr, "0x")

	decodedByte, err := hex.DecodeString(sanitizedByteStr)
	if err != nil {
		handle(err)
	}

	b := decodedByte[0]

	valuesDecode := []bool{
		(b&0x80)>>7 == 1,
		(b&0x40)>>6 == 1,
		(b&0x20)>>5 == 1,
		(b&0x10)>>4 == 1,
		(b&0x8)>>3 == 1,
		(b&0x4)>>2 == 1,
		(b&0x2)>>1 == 1,
		(b&0x1)>>0 == 1,
	}

	fmt.Printf("%v\n", valuesDecode)
	return valuesDecode

	// side effect!
	//fmt.Printf("%t - %t - %t - %t --- %t - %t - %t - %t\n",
	//	oscmanager.SetSlotEnable(8, (b&0x80)>>7 == 1),
	//	oscmanager.SetSlotEnable(7, (b&0x40)>>6 == 1),
	//	oscmanager.SetSlotEnable(6, (b&0x20)>>5 == 1),
	//	oscmanager.SetSlotEnable(5, (b&0x10)>>4 == 1),
	//	oscmanager.SetSlotEnable(4, (b&0x8)>>3 == 1),
	//	oscmanager.SetSlotEnable(3, (b&0x4)>>2 == 1),
	//	oscmanager.SetSlotEnable(2, (b&0x2)>>1 == 1),
	//	oscmanager.SetSlotEnable(1, (b&0x1)>>0 == 1),
	//)

}
