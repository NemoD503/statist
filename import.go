package main

import (
	"fmt"
	"os/exec"
	"log"
)

type Docker struct {
	container string
}

func (i *Docker) Import() (*Stats, error) {
	stats := &Stats{}
	cmd := exec.Command("docker", "exec", i.container, "curl", "http://localhost:2398/stats")
	out, err := cmd.Output()
	if err != nil {
		log.Println("Error on docker command:", err, out)
		return stats, fmt.Errorf("error on docker command: %v", err)
	}

	err = stats.UnmarshalText(out)

	if	err != nil {
		log.Println("Error on unmurshalling:", err)
		return stats, fmt.Errorf("error on unmurshalling: %v", err)
	}

	return stats, nil
}