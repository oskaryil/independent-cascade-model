package multigraph

import (
	"fmt"
	"time"
)

func (g *Graph) DiffuseInformation(seed []int64, diffusionCase string) map[int64]time.Time {
	g.lock.RLock()
	defer g.lock.RUnlock()
	informedNodes := make(map[int64]time.Time)

	for _, nodeID := range seed {
		informedNodes[nodeID] = time.Time{}
	}

	// fmt.Println(informedNodes)

	needsUpdate := true

	for needsUpdate {
		needsUpdate = false
		lines := g.AdjacentEdges(informedNodes)
		for i := range lines {
			u := lines[i].From()
			v := lines[i].To()
			dt := lines[i].DiffusionTime()

			if _, exists := informedNodes[u.ID()]; exists {
				if informedNodes[u.ID()].IsZero() || dt.After(informedNodes[u.ID()]) {
					if _, exists := informedNodes[v.ID()]; !exists {
						informedNodes[v.ID()] = dt
						needsUpdate = true
					} else {
						if !informedNodes[v.ID()].IsZero() && dt.Before(informedNodes[v.ID()]) {
							informedNodes[v.ID()] = dt
							needsUpdate = true
						}
					}
				}
			}

			if _, exists := informedNodes[v.ID()]; exists {
				if informedNodes[v.ID()].IsZero() || dt.After(informedNodes[v.ID()]) {
					if _, exists := informedNodes[u.ID()]; !exists {
						informedNodes[u.ID()] = dt
						needsUpdate = true
					} else {
						if !informedNodes[u.ID()].IsZero() && dt.Before(informedNodes[u.ID()]) {
							informedNodes[u.ID()] = dt
							needsUpdate = true
						}
					}
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
