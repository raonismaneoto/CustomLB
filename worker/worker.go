package main

const (
	NodeEndpoint = "localhost:8081"
)

func main() {
	api := Api{}
	api.Start("8081")
}

type Worker struct {
}

func (w *Worker) join() {

}

func (w *Worker) update() {

}

func (w *Worker) remove() {

}
