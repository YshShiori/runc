// +build linux,!gccgo

package nsenter

/*
#cgo CFLAGS: -Wall
extern void nsexec();
void __attribute__((constructor)) init(void) {
	nsexec();
}
*/
import "C"

// 这里将init函数设置为了__attribute__((constructor)), 使得它会在Go runtime.main之前就被执行

