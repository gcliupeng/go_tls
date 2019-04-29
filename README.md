# go的tls库
golang并没有提供tls相关的接口，使得我们在共享数据时不得不通过传递context等方法实现。这里实现了一个golang的tls，在建立协程时调用go_tls.Set_ctx(V),然后在本协程的任何函数里都可以通过go_tls.Get_ctx()获取数据。
## 实现原理
1. 首先获取golang协程数据结构中的goid字段，获取方法见go_tls.go文件中的GetgoId方法和getg.s文件中的getg方法。
2. 本库包含一个全局的sync.Map，Set_ctx、Get_ctx都是对sync.Map的简单封装
## 使用实例
    package main
    import "fmt"
    import "go_tls"
    
    
    func main() {
    	go_tls.Set_ctx(100)
    	v,ok :=go_tls.Get_ctx()
    	if !ok{
    		panic("error")
    	}
    	if vv,ok:= v.(int);ok{
    		fmt.Println(vv)
    	}
    	go func (n int) {
    		defer func () {
    			go_tls.Del_ctx()
    		}()
    		tls.Set_ctx(n)
    		v,ok :=go_tls.Get_ctx()
    		if !ok{
    			panic("error")
    		}
    		if vv,ok:= v.(int);ok{
    			fmt.Println(vv)
    		}
    	}(102)
    }
    
## 	todo
linux x86系统运行良好，其他环境尚未测试
