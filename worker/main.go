package main

func main() {
	api := Api{
		server: nil,
		worker: Worker{},
	}
	api.Start("8081")
}
