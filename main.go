package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/alexflint/go-arg"
	"golang.org/x/crypto/ssh/terminal"
)

func stringToFloat64(strList []string) []float64 {
	floatList := make([]float64, len(strList))
	for index, str := range strList {
		floatList[index], _ = strconv.ParseFloat(str, 64)
	}
	return floatList
}

func getMaximum(list []float64) float64 {
	return list[len(list)-1]
}

func getMinimum(list []float64) float64 {
	return list[0]
}

func getAverage(list []float64) float64 {
	sum := 0.0
	for _, num := range list {
		sum += num
	}
	return sum / float64(len(list))
}

func getMedian(list []float64) float64 {
	return getPercentile(list, 50)
}

func getPercentile(list []float64, n int) float64 {
	if n == 0 {
		return list[0]
	}
	k := len(list)*n/100 - 1
	return list[k]
}

func main() {
	var reader io.Reader
	var args struct {
		File string `arg:"positional"`
		Min  bool
		Max  bool
		Avg  bool
		Med  bool
		Per  []int
	}

	arg.MustParse(&args)
	isTerminal := terminal.IsTerminal(int(os.Stdin.Fd()))

	if isTerminal {
		if len(args.File) == 0 {
			fmt.Println("[ERROR] NO INPUT FILE")
			os.Exit(1)
		}
		file, _ := os.OpenFile(args.File, os.O_RDONLY, 0666)
		defer file.Close()
		reader = file
	} else {
		reader = os.Stdin
	}

	input, _ := ioutil.ReadAll(reader)

	trimed := strings.Trim(string(input), "\n")

	if len(trimed) == 0 {
		fmt.Println("[ERROR] INPUT FILE IS INVALID")
		os.Exit(1)
	}

	stringList := strings.Split(trimed, "\n")
	floatList := stringToFloat64(stringList)

	sort.Float64s(floatList)

	if args.Min || args.Max || args.Avg || args.Med || len(args.Per) != 0 {
		// オプション指定あり -> 指定ありオプションのみ出力
		if args.Min {
			fmt.Printf("Min %6.6f\n", getMinimum(floatList))
		}
		if args.Max {
			fmt.Printf("Max %6.6f\n", getMaximum(floatList))
		}
		if args.Avg {
			fmt.Printf("Avg %6.6f\n", getAverage(floatList))
		}
		if args.Med {
			fmt.Printf("Med %6.6f\n", getMedian(floatList))
		}
		for _, n := range args.Per {
			fmt.Printf("%2d%% %6.6f\n", n, getPercentile(floatList, n))
		}
	} else {
		// オプション指定なし=デフォルト -> 全オプション出力
		fmt.Printf("Min %6.6f\n", getMinimum(floatList))
		fmt.Printf("Max %6.6f\n", getMaximum(floatList))
		fmt.Printf("Avg %6.6f\n", getAverage(floatList))
		fmt.Printf("Med %6.6f\n", getMedian(floatList))
		for _, n := range []int{90, 95, 99} {
			fmt.Printf("%2d%% %6.6f\n", n, getPercentile(floatList, n))
		}
	}
}
