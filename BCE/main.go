package slice

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	fmt.Print(os.TempDir())

}

//go:noinline
func xy(a []string) {
	x := a[0]
	y := a[1]
	_ = x + y
}

//go:noinline
func yx_faster(a []string) {
	if len(a) > 0 {
		x := a[1]
		y := a[0]
		_ = x + y
	}
}

//go:noinline
func sliceUniqueStd(ss []string) []string {
	seen := make(map[string]bool, len(ss))
	i := 0
	for _, v := range ss {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = true
		ss[i] = v

		i++
	}
	return ss[:i]
}

var sourcePool = sync.Pool{
	New: func() interface{} {
		return make(map[string]bool)
	},
}

//go:noinline
func getSource() map[string]bool {
	r := sourcePool.Get().(map[string]bool)
	return r
}
//go:noinline
func putSource(r map[string]bool) {
	sourcePool.Put(r)
}

//go:noinline
func sliceUniqueUpdated(ss []string) []string {
	seen := make(map[string]bool, len(ss))
	//seen := getSource()
	ii := 0

	for i := 0; i < len(ss); i++ {
		if _, ok := seen[ss[i]]; ok {
			continue
		}

		seen[ss[i]] = true
		if ii > 0 && len(ss) > ii {
			ss[ii] = ss[i]
		}

		ii++
	}

	//putSource(seen)

	if len(ss) > ii && ii > 0 {
		return ss[:ii]
	}

	return nil
}
