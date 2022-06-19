package fuzzer

func Interrupted() bool {
	if runTime > 2 {
		return true
	} else {
		return false
	}
}
