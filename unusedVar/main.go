package main 

func FakeUse(unusedVar any) {
	_= unusedVar
}
func main() {
	unusedVar := 42
	FakeUse(unusedVar)
}
