package utils

// CheckError takes an error type and panics if its an error
func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}
