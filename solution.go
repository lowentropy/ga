package ga

type Solution interface {
	Evaluate()
	Fitness() float64
	Mutate()
	Combine(Solution) (Solution, Solution)
	Copy() Solution
}
