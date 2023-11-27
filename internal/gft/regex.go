package gft

func Regex(testNames []string) string {
	return "^Test(ShouldListTestNamesFromFiles|ShouldReturnErrorIfFileIsNotFound)$"
}
