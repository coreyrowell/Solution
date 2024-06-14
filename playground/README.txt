Taking heuristics such as traffic, time load/unload, etc. into consideration - the problem becomes so complex there may be no one best solution.

`problemSet10Loads.png` - plots the pickup->dropoff for all loads in problem 10 txt file
`problem18Midpoints.png` - plots the midpoints of each load in problem 18 txt file

Visually, there may be an opportunity for better algorithm that clusters data, perhaps on midpoints and routes drivers to clusters to optimize a driver receiving the most possible loads.

`myCode/load/load.go` struct:
is easily extensible to add new properties to aid in clustering:
1) drive time between pickup/dropoff
2) return time between dropoff/depot

Such an algorithm would be easily implemented using the `LoadAssignmentStrategy` interface