package slice

//func main() {
//
//}

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

func sliceUniqueUpdated(ss []string) []string {
	seen := make(map[string]bool, len(ss))
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

	if len(ss) > ii && ii > 0 {
		return ss[:ii]
	}

	return nil
}
