package main

func main() {
	companyUUID := ""
	groupUUID := ""

	if companyUUID != "" && groupUUID != "" {
		println("done")
	} else {
		if groupUUID != "" {
			println("done2")
		} else if companyUUID != "" {
			println("done3")
		} else {
			println("done4")
		}
	}

}
