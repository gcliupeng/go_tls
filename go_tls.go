package go_tls
// import "fmt"
import "unsafe"
import "sync"
type stack struct {
	lo uintptr
	hi uintptr
}

type gobuf struct {
	sp   uintptr
	pc   uintptr
	g    uintptr
	ctxt unsafe.Pointer // this has to be a pointer so that gc scans it
	ret  uint64
	lr   uintptr
	bp   uintptr // for GOEXPERIMENT=framepointer
}

type g struct {
	stack       stack   
	stackguard0 uintptr 
	stackguard1 uintptr 

	_panic         uintptr 
	_defer         uintptr 
	m              uintptr      
	sched          gobuf
	syscallsp      uintptr        
	syscallpc      uintptr        
	stktopsp       uintptr        
	param          unsafe.Pointer 
	atomicstatus   uint32
	stackLock      uint32 
	goid           int64
}

func getg()(uintptr)

func GetgoId() int64{
	gp:=getg()
	gg := (*g)(unsafe.Pointer(gp))
	return gg.goid
}
var smap sync.Map

func Set_ctx(p interface{}) {
	id:= GetgoId()
	smap.Store(id,p)
}
func Get_ctx() (v interface{} , ok bool) {
	id:= GetgoId()
	return smap.Load(id)
}
func Del_ctx() {
	id:= GetgoId()
	smap.Delete(id)
}