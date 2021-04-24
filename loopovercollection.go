package main

import ()

func loopovercollection() {
	slice := []int{1,2,3}
	
	for i:=0;i<len(slice);i++{
	println(slice[i])
	}
	
	for i, v := range slice {
		println(i,v)
	}
	
	wellKnownPorts := map[string]int{"https": 443, "http": 80}
	
	for _,v := range wellKnownPorts{
		println(v)
	}
	
	for i,v := range wellKnownPorts{
		println(i,v)
	}
}
