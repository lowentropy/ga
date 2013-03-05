package ga

import (
	"math/rand"
)

type Ga struct {
	population []Solution
	best       Solution
	score      float64
	generation int
	pMutate    float64
	pCrossover float64
	bouts      int
}

func New(size int) (ga *Ga) {
	population := make([]Solution, size)
	ga = &Ga{population, nil, -999999, 0, 0.01, 0.1, 3}
	ga.Evaluate()
	return
}

func (ga *Ga) Step() Solution {
	ga.Mutate()
	ga.Select()
	ga.Combine()
	ga.Evaluate()
	ga.generation++
	return ga.best
}

func (ga *Ga) Evaluate() {
	eachSolution(ga, func(solution Solution) {
		solution.Evaluate()
		if solution.Fitness() > ga.score {
			ga.score = solution.Fitness()
			ga.best = solution
		}
	})
}

func (ga *Ga) Mutate() {
	eachSolution(ga, func(solution Solution) {
		if rand.Float64() < ga.pMutate {
			solution.Mutate()
		}
	})
}

func (ga *Ga) Select() {
	next := make([]Solution, size(ga))
	upto(size(ga), 1, func(i int) {
		best := randomSolution(ga)
		for j := 0; j < ga.bouts; j++ {
			other := randomSolution(ga)
			if best.Fitness() < other.Fitness() {
				best = other
			}
		}
		next[i] = best.Copy()
	})
	ga.population = next
}

func (ga *Ga) Combine() {
	upto(size(ga), 2, func(i int) {
		if rand.Float64() < ga.pCrossover {
			p1, p2 := ga.population[i], ga.population[i+1]
			c1, c2 := p1.Combine(p2)
			ga.population[i], ga.population[i+1] = c1, c2
		}
	})
}

func upto(n, d int, f func(int)) {
	ch := make(chan int, n/d)
	for i := 0; i < n; i += d {
		go func() {
			f(i)
			ch <- 1
		}()
	}
	for i := 0; i < n; i += d {
		<-ch
	}
}

func eachSolution(ga *Ga, f func(Solution)) {
	ch := make(chan int, size(ga))
	for _, solution := range ga.population {
		go func() {
			f(solution)
			ch <- 1
		}()
	}
	for i := 0; i < size(ga); i++ {
		<-ch
	}
}

func randomSolution(ga *Ga) Solution {
	return ga.population[rand.Intn(size(ga))]
}

func size(ga *Ga) int {
	return len(ga.population)
}
