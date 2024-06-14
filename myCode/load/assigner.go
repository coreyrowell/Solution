package load


type LoadAssignmentStrategy interface {
    AssignLoadsToDriver(loads []Load) [][]int
}

type LoadAssigner struct {
    strategy LoadAssignmentStrategy
}

func (la *LoadAssigner) SetStrategy(strategy LoadAssignmentStrategy) {
    la.strategy = strategy
}

func (la *LoadAssigner) AssignLoads(loads []Load) [][]int {
    return la.strategy.AssignLoadsToDriver(loads)
}