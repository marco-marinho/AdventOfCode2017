package main

import (
	"aoc/helpers"
	"fmt"
	"math"
	"sort"
	"strings"
)

type Particle struct {
	pos []int
	vel []int
	acc []int
}

func main() {
	particles := getParticles()
	for i := 0; i < 1000; i++ {
		step(particles)
	}
	fmt.Println("Task 01:", getClosest(particles))

	particles = getParticles()
	for i := 0; i < 100; i++ {
		step(particles)
		particles = removeCollisions(particles)
	}
	fmt.Println("Task 02:", len(particles))
}

func getParticles() []Particle {
	data := helpers.GetLines("Data/Day20.txt")
	particles := make([]Particle, len(data))
	for idx, line := range data {
		parts := strings.Split(line, ", ")
		posRaw := strings.ReplaceAll(parts[0], "p=<", "")
		posRaw = strings.ReplaceAll(posRaw, ">", "")
		pos := helpers.StringToInts(posRaw, ",")
		velRaw := strings.ReplaceAll(parts[1], "v=<", "")
		velRaw = strings.ReplaceAll(velRaw, ">", "")
		vel := helpers.StringToInts(velRaw, ",")
		accRaw := strings.ReplaceAll(parts[2], "a=<", "")
		accRaw = strings.ReplaceAll(accRaw, ">", "")
		acc := helpers.StringToInts(accRaw, ",")
		particles[idx] = Particle{pos, vel, acc}
	}
	return particles
}

func getClosest(particles []Particle) int {
	min := math.MaxInt
	idxMin := -1
	for idx, particle := range particles {
		dist := helpers.AbsInt(particle.pos[0]) + helpers.AbsInt(particle.pos[1]) + helpers.AbsInt(particle.pos[2])
		if dist < min {
			min = dist
			idxMin = idx
		}
	}
	return idxMin
}

func step(particles []Particle) {
	for _, particle := range particles {
		particle.vel[0] += particle.acc[0]
		particle.vel[1] += particle.acc[1]
		particle.vel[2] += particle.acc[2]
		particle.pos[0] += particle.vel[0]
		particle.pos[1] += particle.vel[1]
		particle.pos[2] += particle.vel[2]
	}
}

func removeCollisions(particles []Particle) []Particle {
	toRemove := make(map[int]bool)
	for idxOuter := 0; idxOuter < len(particles); idxOuter++ {
		for idxInner := idxOuter + 1; idxInner < len(particles); idxInner++ {
			particleOuter := particles[idxOuter]
			particleInner := particles[idxInner]
			if samePos(particleOuter, particleInner) {
				toRemove[idxOuter] = true
				toRemove[idxInner] = true
			}
		}
	}
	toRemoveSlice := make([]int, len(toRemove))
	idx := 0
	for key := range toRemove {
		toRemoveSlice[idx] = key
		idx++
	}
	sort.Slice(toRemoveSlice, func(i, j int) bool {
		return toRemoveSlice[i] > toRemoveSlice[j]
	})
	for _, removeIdx := range toRemoveSlice {
		particles = remove(particles, removeIdx)
	}
	return particles
}

func samePos(particle1 Particle, particle2 Particle) bool {
	return particle1.pos[0] == particle2.pos[0] && particle1.pos[1] == particle2.pos[1] && particle1.pos[2] == particle2.pos[2]
}

func remove(slice []Particle, s int) []Particle {
	return append(slice[:s], slice[s+1:]...)
}
