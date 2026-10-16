package abstractions

type IWebProxy interface {
	GetProxy(destination string) // method to get the proxy
	IsBypassed(host string) bool // method to check if a host is bypassed
}
