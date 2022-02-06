package main

import "Winx/wnet"

func main() {
	s := wnet.NewServer("[Winx v0.2]")
	s.Serve()
}
