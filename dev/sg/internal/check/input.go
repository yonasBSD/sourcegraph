package check

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/sourcegraph/sourcegraph/dev/sg/internal/std"
	"github.com/sourcegraph/sourcegraph/lib/output"
)

const (
	FixAll  = 0
	FixQuit = 99
)

func getChoice(in io.Reader, out *std.Output, choices map[int]string) (int, error) {
	// maps are not guaranteed to be in the same order, so
	// we need to sort the choices so that we can print them in the same order every time.
	choiceIdx := make([]int, len(choices))
	idx := 0
	for num := range choices {
		choiceIdx[idx] = num
		idx++
	}
	sort.Ints(choiceIdx)

	for {
		out.Write("")
		out.WriteNoticef("What do you want to do?")

		for _, num := range choiceIdx {
			desc := choices[num]
			out.Writef("%s[%2d]%s: %s", output.StyleBold, num, output.StyleReset, desc)
		}

		out.Promptf("Enter choice:")
		var s int
		if _, err := fmt.Fscan(in, &s); err != nil {
			return 0, err
		}

		if _, ok := choices[s]; ok {
			return s, nil
		}
		out.WriteFailuref("Invalid choice")
	}
}

func askWhatToFix(in io.Reader, out *std.Output, numbers []int) (int, error) {
	var strs []string
	var idx = make(map[int]struct{})
	for _, num := range numbers {
		strs = append(strs, fmt.Sprintf("%d", num+1))
		idx[num+1] = struct{}{}
	}

	for {
		out.Promptf("[%s]:", strings.Join(strs, ","))
		var num int
		_, err := fmt.Fscan(in, &num)
		if err != nil {
			return 0, err
		}

		if _, ok := idx[num]; ok {
			return num - 1, nil
		}
		out.Writef("%d is an invalid choice :( Let's try again?\n", num)
	}
}

func waitForReturn(in io.Reader) {
	fmt.Fscanln(in)
}
