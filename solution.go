package ga

type Solution interface {
	EvaluateFitness()
	Fitness() float64
	Mutate()
	Copy() Solution
	Combine(*Solution) (Solution, Solution)
}
