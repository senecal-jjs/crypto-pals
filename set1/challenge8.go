package set1

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/senecal-jjs/crypto-pals/util"
)

func Challenge8() {
	file, err := os.ReadFile("set1/input8.txt")
	util.PanicOnErr(err)

	inputs := strings.Split(string(file), "\n")

	bestScore := -1
	bestInput := ""

	for _, c := range inputs {
		iBytes, err := hex.DecodeString(c)
		util.PanicOnErr(err)

		chunks := chunkByteArray(iBytes, 16)
		repeats := countDuplicates(chunks)

		if repeats > bestScore {
			bestScore = repeats
			bestInput = c
		}
	}

	fmt.Printf("Repeats: %d\nHex String: %q\n", bestScore, bestInput)
}

func countDuplicates(bufs [][]byte) int {
	repeats := 0

	for i, chunk1 := range bufs {
		for j, chunk2 := range bufs {
			if i != j {
				if bytes.Equal(chunk1, chunk2) {
					repeats++
				}
			}
		}
	}

	return repeats
}
