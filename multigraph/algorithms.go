package multigraph

import (
	"fmt"
	"time"
)

func (g *Graph) DiffuseInformation(seed []int64, beta float64, diffusionCase string) map[int64]time.Time {
	g.lock.Lock()
	defer g.lock.Unlock()
	informedNodes := make(map[int64]time.Time)
	changed := true

	for _, nodeID := range seed {
		informedNodes[nodeID] = time.Time{}
	}

	for changed {
		lines := g.AdjacentEdges(informedNodes)
		fmt.Println(lines)
		for i := range lines {
			u := lines[i].From()
			v := lines[i].To()
			dt := lines[i].DiffusionTime()

			changed = false

			if informedNodes[u.ID()].IsZero() || informedNodes[u.ID()].Before(dt) {
				if nodeV, exists := informedNodes[v.ID()]; !exists || (dt.Before(nodeV) && !nodeV.IsZero()) {
					fmt.Println(exists, dt)
					fmt.Println(v.ID())
					informedNodes[v.ID()] = dt
					changed = true
				}
			}
		}
	}
	return informedNodes
}
