package main

import (
	"os"
	"testing"
)

func TestProcessFile(t *testing.T) {

}

var expected = Result{
	numBytes:      2421,
	numCharacters: 2421,
	numWords:      351,
	numLines:      8,
}

var flagSet = FlagSet{
	characterFlag: true,
	byteFlag:      true,
	wordFlag:      true,
	lineFlag:      true,
}

func TestProcessFileByRunesOnDefaultFlagSet(t *testing.T) {
	// default: with no flags provided, only byte, word, and line count should be computed
	// arrange | act | assert
	fd, err := os.Open("./testdata/data.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()
	flagSet.characterFlag = false
	res, err := processFileByRunes(fd, flagSet)
	if err != nil {
		t.Fatal(err)
	}
	if res.numCharacters != 0 {
		t.Error("Character flag is false and returned numCharacters is not zero")
	}
	if res.numBytes != expected.numBytes {
		t.Errorf("received numBytes %d different from expected %d", res.numBytes, expected.numBytes)
	}

	if res.numWords != expected.numWords {
		t.Errorf("received numWords %d different from expected %d", res.numWords, expected.numWords)
	}

	if res.numLines != expected.numLines {
		t.Errorf("received numLines %d different from expected %d", res.numLines, expected.numLines)
	}
}
