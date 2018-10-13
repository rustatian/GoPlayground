package slice

import (
	"sync"
)

func main() {

}

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

func getSource() map[string]bool {
	r := sourcePool.Get().(map[string]bool)
	return r
}
func putSource(r map[string]bool) {
	sourcePool.Put(r)
}

func sliceUniqueUpdated(ss []string) []string {
	nSS := ss
	//seen := make(map[string]bool, len(nSS))
	seen := getSource()
	ii := 0

	for i := 0; i < len(nSS); i++ {
		if _, ok := seen[nSS[i]]; ok {
			continue
		}

		seen[nSS[i]] = true
		if ii > 0 && len(nSS) > ii {
			nSS[ii] = nSS[i]
		}

		ii++
	}

	putSource(seen)

	if len(nSS) > ii && ii > 0 {
		return nSS[:ii]
	}

	return nil
}
