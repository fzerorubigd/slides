package main
import (
	πg "grumpy"
	π_os "os"
)
func initModule(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
	var πTemp001 []*πg.Object
	_ = πTemp001
	var πE *πg.BaseException; _ = πE
	for ; πF.State() >= 0; πF.PopCheckpoint() {
		switch πF.State() {
		case 0:
		default: panic("unexpected function state")
		}
		// line 1: print "hello world"
		πF.SetLineno(1)
		πTemp001 = make([]*πg.Object, 1)
		πTemp001[0] = πg.NewStr("hello world").ToObject()
		if πE = πg.Print(πF, πTemp001, true); πE != nil {
			continue
		}
		return nil, nil
	}
	return nil, πE
}
var Code *πg.Code
func main() {
	Code = πg.NewCode("<module>", "./hello.py", nil, 0, initModule)
	π_os.Exit(πg.RunMain(Code))
}
