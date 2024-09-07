package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Service struct {
	Name    string
	Healthy bool
	lock    sync.Mutex
}

func (s *Service) Monitor(quit <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-quit:
			fmt.Printf("Stopping monitoring of service: %s\n", s.Name)
			return
		default:
			s.lock.Lock()
			if s.Healthy {
				fmt.Printf("Service %s is healthy.\n", s.Name)
			} else {
				fmt.Printf("Service %s is down.\n", s.Name)
			}
			s.lock.Unlock()
			time.Sleep(time.Second * 1)
		}
	}

}

// SimulateHealth function randomly changes the health of the service
func (s *Service) SimulateHealth() {
	for {
		s.lock.Lock()
		s.Healthy = rand.Intn(2) == 0
		s.lock.Unlock()
	}
}

func main() {
	service := []Service{
		{Name: "Servive 1", Healthy: true},
		{Name: "Servive 2", Healthy: true},
		{Name: "Servive 3", Healthy: true},
		{Name: "Servive 4", Healthy: true},
	}
	quit := make(chan bool)
	var wg sync.WaitGroup

	for _, service := range service {
		wg.Add(1)
		go service.Monitor(quit, &wg)
		go service.SimulateHealth()
	}

	time.Sleep(time.Second * 10)
	close(quit)
	wg.Wait()
	fmt.Println("All monitoring stopped.")
}
