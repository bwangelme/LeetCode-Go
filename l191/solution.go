package l191

import (
	"log"
	"strconv"
)

func hammingWeight(num uint32) int {
	if num == 0 {
		return 0
	}

	var count = 1

	for {
		num &= (num - 1)
		if num == 0 {
			break
		}
		count += 1

	}
	return count
}

func binStrToUint32(binary string) uint32 {
	num, err := strconv.ParseUint(binary, 2, 32)
	if err != nil {
		log.Fatalln(err)
	}

	return uint32(num)
}
