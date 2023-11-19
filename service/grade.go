// golang.org/x/tools/cmd/godoc
package service

func CheckGrade(score int) string {
	switch {
	case score == 100:
		return "A+"
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

// Hard Function to test
func DatabaseFunction() {

}
