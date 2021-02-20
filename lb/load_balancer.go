package lb

type Node struct {
	id string
	endpoint string
	totalCpu float32 // in mb
	usedCpu float32 // varies from 0 to 1
	totalRam float32 // in mb
	usedRam float32 // varies from 0 to 1
	coef float32 // coefficient calculated based on the node resources. Varies from 0, when the node is unable to receive any request, to 1, when the node is completely free
}

type LoadBalancer struct {
	nodes []Node
}

func (lb *LoadBalancer) Start() {
 go lb.sortNodes()
}

func (lb *LoadBalancer) AddNode() {

}

func (lb *LoadBalancer) ResolveRequest() {

}

func (lb *LoadBalancer) RemoveNode() {

}

func (lb *LoadBalancer) sortNodes() {
	for {

	}
}
