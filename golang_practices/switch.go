package main

type HTTPRequest struct {
	Method string
}

func switch_ex() {
	r := HTTPRequest{Method: "Get"}

	switch r.Method {
	case "Get":
		println("Get")

	case "Post":
		println("Post")

	case "Delete":
		println("Delete")

	default:
		println("default stmt")
	}
}
