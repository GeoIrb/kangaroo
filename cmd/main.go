package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/geoirb/kangaroo/pkg/kangaroo"
	"github.com/geoirb/kangaroo/pkg/utils"
)

const (
	epsilon = 1e-9

	serviceName = "kangaroo"
)

func main() {
	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.WithPrefix(logger, "service", serviceName)

	level.Info(logger).Log("msg", "initializing")

	number := utils.NewNumber(epsilon)
	fabric := kangaroo.NewKangarooFabric(number)

	runtime := func(input string) {
		defer fmt.Printf("Input:\t")
		
		strs := strings.Split(input, " ")
		if len(strs) != 4 {
			level.Error(logger).Log("err", "few parameters")
			return
		}

		x, err := strconv.Atoi(strs[0])
		if err != nil {
			level.Error(logger).Log("msg", "convert parameter", "parameter number", 0, "err", err)
			return
		}

		v, err := strconv.Atoi(strs[1])
		if err != nil {
			level.Error(logger).Log("msg", "convert parameter", "parameter number", 1, "err", err)
			return
		}

		first := fabric(x, v)

		x, err = strconv.Atoi(strs[2])
		if err != nil {
			level.Error(logger).Log("msg", "convert parameter", "parameter number", 2, "err", err)
			return
		}

		v, err = strconv.Atoi(strs[3])
		if err != nil {
			level.Error(logger).Log("msg", "convert parameter", "parameter number", 3, "err", err)
			return
		}

		second := fabric(x, v)

		result := "NO"
		if first.IsIntersect(second) {
			result = "YES"
		}

		fmt.Printf("Output:\t%s\n", result)
	}

	fmt.Printf("Input:\t")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		runtime(scanner.Text())
	}
}
