// Package stack provides the glue between networking protocols and the
// consumers of the networking stack.
//
// For consumers, the only function of interest is New(), everything else is
// provided by the tcpip/public package.
//
// For protocol implementers, RegisterTransportProtocol() and
// RegisterNetworkProtocol() are used to register protocols with the stack,
// which will then be instantiated when consumers interact with the stack.
package stack
