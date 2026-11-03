package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var characterFlag = flag.Bool("m", false, "the file character count")
var byteFlag = flag.Bool("c", false, "the file byte count")
var lineFlag = flag.Bool("l", false, "the file line count")
var wordFlag = flag.Bool("w", false, "the file word count")

type Result struct {
	numLines      int
	numWords      int
	numBytes      int
	numCharacters int
}

type FlagSet struct {
	characterFlag bool
	wordFlag      bool
	lineFlag      bool
	byteFlag      bool
}

func (fs *FlagSet) isDefault() bool {
	return !fs.byteFlag && !fs.characterFlag && !fs.lineFlag && !fs.wordFlag
}

func main() {
	// main.go -c file to output the number of bytes
	// main.go -c -l -w -m file.txt
	// echo "hello world" | main.go
	// output format: numLines, numWords, numCharacters, numBytes
	flag.Parse()

	flagSet := FlagSet{
		characterFlag: *characterFlag,
		byteFlag:      *byteFlag,
		wordFlag:      *wordFlag,
		lineFlag:      *lineFlag,
	}

	if flagSet.isDefault() {
		flagSet = FlagSet{
			byteFlag: true,
			wordFlag: true,
			lineFlag: true,
		}
	}

	// m := "sa   3i   i"
	// r := []rune(m)
	// fmt.Print(r)
	flagArgs := flag.Args()
	var fd *os.File
	var err error
	var filePath string

	if len(flagArgs) > 0 {
		filePath = flagArgs[0]
		fd, err = os.Open(filePath)
		if err != nil {
			fmt.Printf("error: %v", err)
			os.Exit(1)
		}
		defer fd.Close()
	} else {
		fd = os.Stdin
	}
	result, err := processFileByRunes(fd, flagSet)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	printResult(result, flagSet, filePath)
}

func printResult(result *Result, flagSet FlagSet, filePath string) {
	resPrint := make([]string, 0, 10)

	if flagSet.lineFlag {
		resPrint = append(resPrint, strconv.FormatInt(int64(result.numLines), 10))
	}

	if flagSet.wordFlag {
		resPrint = append(resPrint, strconv.FormatInt(int64(result.numWords), 10))
	}

	if flagSet.characterFlag {
		resPrint = append(resPrint, strconv.FormatInt(int64(result.numCharacters), 10))
	}

	if flagSet.byteFlag {
		resPrint = append(resPrint, strconv.FormatInt(int64(result.numBytes), 10))
	}

	resPrint = append(resPrint, filePath)
	fmt.Println(strings.Join(resPrint, " "))
}

func processFile(fd *os.File, flagSet FlagSet) (*Result, error) {
	result := &Result{}

	// buffered reader
	br := bufio.NewReader(fd)
	for {
		line, err := br.ReadString('\n')
		if err != nil {
			// for potential data read last before \n delimiter is reached
			if err == io.EOF {
				if flagSet.byteFlag {
					result.numBytes += getByteCount(line)
				}
				if flagSet.characterFlag {
					result.numCharacters += getCharacterCount(line)
				}
				if flagSet.wordFlag {
					result.numWords += getWordCount(line)
				}
				break
			}
			return nil, err
		}

		if flagSet.byteFlag {
			result.numBytes += getByteCount(line)
		}
		if flagSet.characterFlag {
			result.numCharacters += getCharacterCount(line)
		}
		if flagSet.wordFlag {
			result.numWords += getWordCount(line)
		}

		if flagSet.lineFlag {
			result.numLines++
		}
	}
	return result, nil
}

func processFileByRunes(fd *os.File, flagSet FlagSet) (*Result, error) {
	result := &Result{}

	// buffered reader
	br := bufio.NewReader(fd)
	charInWord := false
	for {
		char, size, err := br.ReadRune()
		if err != nil {
			// before the EOF was encountered, we check if the last char was in word to increment word count
			if err == io.EOF {
				if flagSet.wordFlag && charInWord {
					result.numWords++
				}
				break
			}
			return nil, err
		}

		if flagSet.byteFlag {
			result.numBytes += size
		}
		if flagSet.characterFlag {
			result.numCharacters++
		}

		if unicode.IsSpace(char) {
			// charInWord is set after this IsSpace comparison has happened so, it will still
			// represent the last character before it's changed -> if a space is encountered
			// and the last char was in a word, means we just reached the end of the word
			if flagSet.wordFlag && charInWord {
				result.numWords++
			}

			// check to confirm the unicode space character is the new line
			if flagSet.lineFlag && char == '\n' {
				result.numLines++
			}
			// char is in a space now so, reset
			charInWord = false
		} else {
			charInWord = true
		}
	}
	return result, nil
}

func getWordCount(line string) int {
	return len(strings.Fields(line))
}

func getCharacterCount(line string) int {
	return len([]rune(line))
}

func getByteCount(line string) int {
	return len([]byte(line))
}
