package main

type HTTPRequest struct {
	Method string
}

func swtiches() {
	r := HTTPRequest{Method: "HEAD"}

	switch r.Method {
	case "GET":
		println("GET Request")
	//fallthrough bir sonraki case e de girer
	case "DELETE":
		println("DELETE Request")
	case "POST":
		println("POST Request")
	case "PUT":
		println("PUT Request")
	default:
		println("Unhandled method")
	}
}
