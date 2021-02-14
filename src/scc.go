package simplelib

import "fmt"

type SCCGo interface {
	SimpleClassCallback
	deleteSimpleClassCallback()
}

type sccGo struct {
	SimpleClassCallback
}

func (scc *sccGo) deleteSimpleClassCallback() {
	DeleteDirectorSimpleClassCallback(scc.SimpleClassCallback)
}

type overwrittenMethodsOnSimpleClassCallback struct {
	scc SimpleClassCallback
}

func (om *overwrittenMethodsOnSimpleClassCallback) OnStart()  {
	fmt.Println("Go Function")
}

func NewSCCGo() SCCGo {
	om := &overwrittenMethodsOnSimpleClassCallback{}
	scc := NewDirectorSimpleClassCallback(om)
	om.scc = scc

	scci := &sccGo{SimpleClassCallback: scc}
	return scci
}

func DeleteSCCGo(scc SCCGo) {
	scc.deleteSimpleClassCallback()
}