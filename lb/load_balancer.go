package lb

import (
	"net/http"
	"time"
)

type Node struct {
	id string
	endpoint string
	totalCpu float32 // in mb
	usedCpu float32 // varies from 0 to 1
	totalRam float32 // in mb
	usedRam float32 // varies from 0 to 1
	coef float32 // coefficient calculated based on the node resources. Varies from 0, when the node is unable to receive any request, to 1, when the node is completely free
	lastUpdatedAt *time.Time
}

type Request struct {
	Method string
	ResourcePath string
	Body interface{}
	Headers http.Header
}

type LoadBalancer struct {
	nodes []*Node
}

func (lb *LoadBalancer) Start() {
 go lb.refreshNodes()
}

func (lb *LoadBalancer) AddNode(node *Node) {
	for i, currNode := range lb.nodes {
		if currNode.coef <= node.coef {
			lb.nodes = append(lb.nodes[:i+1], lb.nodes[i:]...)
			lb.nodes[i] = node
			return
		}
	}
}

func (lb *LoadBalancer) ResolveRequest(incomingReq *Request) (*HttpResponse, error) {
	nodeEndpoint := lb.nodes[0].endpoint
	return SendRequest(nodeEndpoint, incomingReq)
}

func (lb *LoadBalancer) RemoveNode(id string) {

}

func (lb *LoadBalancer) UpdateNode(id string) {

}

func (lb *LoadBalancer) refreshNodes() {
	for {

	}
}

func calculateCoef(node *Node) {

}
