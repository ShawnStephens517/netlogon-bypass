/*package main

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (

	netlogon32DLL = windows.NewLazyDLL("netlogon.dll")
	procIdentityInfo = netlogon32DLL.NewProc("_NETLOGON_IDENTITY_INFO")
)

func main() {
	fmt.Println("[+] Attempting to force NTLMv1 using the specific struct identified by SilverFort")
	fmt.Println("https://www.silverfort.com/blog/ntlmv1-bypass-in-active-directory-technical-deep-dive/")
	procIdentityInfo.Call(doaminName, 0x00010000, UserName, Workstation)

	#Parameter controlflag needed 0x00010000
}
*/