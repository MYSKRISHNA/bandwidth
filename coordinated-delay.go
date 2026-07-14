import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type CommunicationNode struct {
	ID                  int
	TransmissionDelay   int
	CongestionLevel     int
	SynchronizationRate int
	PacketFlow          int
	CommunicationState  bool
}

func monitorNode(
	node *CommunicationNode,
	wg *sync.WaitGroup,
) {

	defer wg.Done()

	time.Sleep(
		time.Millisecond * 300,
	)

	node.TransmissionDelay =
		rand.Intn(250)

	node.CongestionLevel =
		rand.Intn(120)

	node.SynchronizationRate =
		rand.Intn(100)

	node.PacketFlow =
		rand.Intn(500)

	if node.TransmissionDelay > 140 ||
		node.CongestionLevel > 70 ||
		node.SynchronizationRate < 40 {

		node.CommunicationState = false

	} else {

		node.CommunicationState = true
	}
}

func stabilizeNode(
	node *CommunicationNode,
) {

	time.Sleep(
		time.Millisecond * 200,
	)

	node.TransmissionDelay -= 30
	node.CongestionLevel -= 20
	node.SynchronizationRate += 25
	node.PacketFlow += 40

	if node.TransmissionDelay < 140 &&
		node.CongestionLevel < 70 &&
		node.SynchronizationRate > 40 {

		node.CommunicationState = true
	}
}

func executeCluster(
	cluster string,
	nodes int,
) {

	var wg sync.WaitGroup

	services := make(
		[]CommunicationNode,
		nodes,
	)

	for i := 0; i < nodes; i++ {

		services[i] = CommunicationNode{
			ID:                  i + 1,
			TransmissionDelay:   0,
			CongestionLevel:     0,
			SynchronizationRate: 100,
			PacketFlow:          0,
			CommunicationState:  true,
		}
	}

	for i := range services {

		wg.Add(1)

		go monitorNode(
			&services[i],
			&wg,
		)
	}

	wg.Wait()

	unstable := 0

	for i := range services {

		if !services[i].
			CommunicationState {

			stabilizeNode(
				&services[i],
			)

			unstable++
		}
	}

	fmt.Println(
		"Cluster:",
		cluster,
		" Nodes:",
		nodes,
		" Unstable:",
		unstable,
		" Communication Stabilized",
	)
}

func main() {

	rand.Seed(
		time.Now().UnixNano(),
	)

	clusterConfigurations := []struct {
		Name      string
		NodeCount int
	}{
		{"Cluster-A", 3},
		{"Cluster-B", 5},
		{"Cluster-C", 7},
		{"Cluster-D", 9},
		{"Cluster-E", 11},
	}

	for _, config := range
		clusterConfigurations {

		executeCluster(
			config.Name,
			config.NodeCount,
		)
	}

	fmt.Println(
		"Runtime Communication Coordination Completed",
	)
}
