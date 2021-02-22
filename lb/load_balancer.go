package lb

import (
	"encoding/json"
	"errors"
	"fmt"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"log"
	"net/http"
	"sort"
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

func (lb *LoadBalancer) AddNode(id string, node *Node) {
	node.id = id
	lb.nodes = append(lb.nodes, node)
}

func (lb *LoadBalancer) ResolveRequest(incomingReq *Request) (*HttpResponse, error) {
	if len(lb.nodes) == 0 {
		return nil, errors.New("No node available")
	}

	nodeEndpoint := lb.nodes[0].endpoint
	return SendRequest(nodeEndpoint, incomingReq)
}

func (lb *LoadBalancer) RemoveNode(id string) {
	for i, node := range lb.nodes {
		if node.id == id {
			lb.nodes = append(lb.nodes[:i], lb.nodes[i+1:]...)
			fmt.Printf("Node %s removed", id)
			return
		}
	}
}

func (lb *LoadBalancer) UpdateNode(id string, patch jsonpatch.Patch) error {
	for _, node := range lb.nodes {
		if node.id == id {
			parsedNode, err := json.Marshal(node)
			if err != nil {
				log.Printf("unable to parse node %s to byte slice. ", id, err.Error())
			}

			modifiedNode, err := patch.Apply(parsedNode)
			if err != nil {
				log.Printf("unable to apply patch. ", err.Error())
			}

			err = json.Unmarshal(modifiedNode, &node)
			if err != nil {
				log.Printf("unable to unmarshal modified node", err.Error())
			}

			return nil
		}
	}
	return errors.New("node not found")
}

func (lb *LoadBalancer) refreshNodes() {
	for {
		for _, node := range lb.nodes {
			node.coef = calculateCoef(node)
		}

		sort.Slice(lb.nodes, func(i, j int) bool {
			return lb.nodes[i].coef >= lb.nodes[j].coef
		})

		time.Sleep(60 * time.Second)
	}
}

func calculateCoef(node *Node) float32 {
	if node.usedCpu == 0 || node.usedRam == 0 {
		return 1
	}
	coef := (1/(node.usedCpu * node.usedRam))*0.1
	return coef
}
