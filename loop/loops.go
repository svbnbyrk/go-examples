package main

func main() {
	//var i int şeklinde tanımlarsam döngü dışında da int değerini kullanabilirim
	for i := 0; i < 6; i++ {
		if i == 4 {
			break
		}
		if i == 3 {
			continue
		}
	}
	//sonsuz döngü
	for {

	}
}
