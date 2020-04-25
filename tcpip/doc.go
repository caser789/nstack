// Package tcpip provides the interfaces and related types that users of the
// tcpip stack will use in order to create endpoints used to send and receive
// data over the network stack.
//
// The starting point is the creation and configuration of a stack.
// 1. A stack can be created by calling the New() function of the tcpip/stack package;
// 2. configuring a stack involves creating NICs (via calls to Stack.CreateNIC()),
// 3.  adding network addresses (via calls to Stack.AddAddress()), and
// 4. setting a route table (via a call to Stack.SetRouteTable()).
//
// Once a stack is configured, endpoints can be created by calling
// Stack.NewEndpoint(). Such endpoints can be used to send/receive data, connect
// to peers, listen for connections, accept connections, etc., depending on the
// transport protocol selected
package tcpip
