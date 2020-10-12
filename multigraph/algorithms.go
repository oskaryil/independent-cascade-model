package multigraph

import (
	"fmt"
	"time"
)

func (g *Graph) DiffuseInformation(seed []int64, beta float64, diffusionCase string) map[int64]time.Time {
	g.lock.RLock()
	defer g.lock.RUnlock()
	informedNodes := make(map[int64]time.Time)
	changed := true

	for _, nodeID := range seed {
		informedNodes[nodeID] = time.Time{}
	}

	// fmt.Println(informedNodes)

	for changed {
		lines := g.AdjacentEdges(informedNodes)
		// fmt.Println(lines)
		// fmt.Println(lines)
		// fmt.Println(lines)
		for i := range lines {
			u := lines[i].From()
			v := lines[i].To()
			dt := lines[i].DiffusionTime()

			changed = false

			if informedNodes[u.ID()].IsZero() || informedNodes[u.ID()].Before(dt) {
				if nodeV, exists := informedNodes[v.ID()]; !exists || (dt.Before(nodeV) && !nodeV.IsZero()) {
					// fmt.Println(exists, dt)
					// fmt.Println(v.ID())
					informedNodes[v.ID()] = dt
					changed = true
				}
			}
		}
	}
	return informedNodes
}

func (g *Graph) DiffuseInformationSimple(seed []int64, diffusionCase string) map[int64]int64 {
	g.lock.RLock()
	defer g.lock.RUnlock()
	informedNodes := make(map[int64]int64)
	for _, nodeID := range seed {
		informedNodes[nodeID] = 0
	}
	needsUpdate := true

	for needsUpdate {
		needsUpdate = false
		lines := g.AdjacentEdgesSimple(informedNodes)
		fmt.Println("iteration")
		for i := range lines {
			u := lines[i].From()
			v := lines[i].To()
			dn := lines[i].DiffusionNumber()

			if _, exists := informedNodes[u.ID()]; exists {
				if informedNodes[u.ID()] == 0 || dn > informedNodes[u.ID()] {
					if _, exists := informedNodes[v.ID()]; !exists {
						informedNodes[v.ID()] = dn
						needsUpdate = true
					} else {
						if informedNodes[v.ID()] != 0 && dn < informedNodes[v.ID()] {
							informedNodes[v.ID()] = dn
							needsUpdate = true
						}
					}
				}
			}

			if _, exists := informedNodes[v.ID()]; exists {
				if informedNodes[v.ID()] == 0 || dn > informedNodes[v.ID()] {
					if _, exists := informedNodes[u.ID()]; !exists {
						informedNodes[u.ID()] = dn
						needsUpdate = true
					} else {
						if informedNodes[u.ID()] != 0 && dn < informedNodes[u.ID()] {
							informedNodes[u.ID()] = dn
							needsUpdate = true
						}
					}
				}
			}
		}
	}
	return informedNodes
}
