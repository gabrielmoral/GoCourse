package cache

import (
	"fmt"
	"sync"

	"github.com/pabloos/http/greet"
)

type GreetingCached struct {
	Name      string
	Locations []string
}

func (c GreetingCached) Print() string {
	return fmt.Sprintf("From cache - %+v", c)
}

type Greetings struct {
	List map[string][]string
	sync.Mutex
}

func (greetings Greetings) Add(g greet.Greet) (cached *GreetingCached, found bool) {
	locations, ok := greetings.List[g.Name]

	if ok {
		if findLocation(locations, g.Location) {
			return &GreetingCached{Name: g.Name, Locations: greetings.List[g.Name]}, true
		} else {
			greetings.Lock()
			greetings.List[g.Name] = append(greetings.List[g.Name], g.Location)
			greetings.Unlock()
		}
	} else {
		greetings.Lock()
		greetings.List[g.Name] = []string{g.Location}
		greetings.Unlock()
	}
	return nil, false
}

func findLocation(locations []string, location string) bool {
	found := false
	for _, l := range locations {
		if l == location {
			found = true
			break
		}
	}
	return found
}
