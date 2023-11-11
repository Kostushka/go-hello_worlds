package main

import "fmt"

func main() {
	type Flags uint
	const (
		FlagUp Flags = 1 << iota
		FlagBroadcast
		FlagLoopback
		FlagPointToPoint
		FlagMulticast	
	)

	fmt.Printf("%b %b %b %b %b\n", FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)

	const (
		_ = 1 << (10 * iota)
		KiB
		MiB
		GiB
		TiB
		PiB
		EiB
		ZiB
		YiB
	)

	fmt.Printf("%d %d %d\n", KiB, MiB, GiB)
}
