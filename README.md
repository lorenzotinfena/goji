# Data structures and algorithms Golang library
![](goji-image.jpg)

Browser the repo here: https://pkg.go.dev/github.com/lorenzotinfena/goji

## Showcase
- Graph
- Dijkstra
- DFS and BFS
- Heap
- Hashset
- MultiHashSet
- HashMap
- MultiHashMap
- Segment tree
- Lazy segment tree
- HAMT
- Bitset
- LinkedList
- Queue
- Stack
- Tuple
- DP
- Diophantine equations
- Polynomials
- Binary exponentation
- GCD and LCM
- Binary search
- Selection sort
- Mo's algorithm
- Sqrt decomposition
- Radix sort (LSD and MSD)
- Prime generator in O(n)
- Fibonacci heap
- Graham Scan
- Chan's algorithm
- Manacher's algorithm
- Hook length
## Coding guidelines
- When write generic code there are 2 approaches:
    1. OOP: Constraint a type with a method
    2. Procedural: Pass the method directly as a parameter

    If the expected passed type is:
    - user defined: Use procedural
    - internal (defined in this module): Use OOP
    The reason is that an OOP approach is more elegant but less generic (hard to adapt when using multiple libraries with the same type), and procedural method is less elegant, but more generic.

## Assumptions
1. Every pointer receiver method assume that the pointer is not nil
2. All assumption (if any) are specified in each function
3. Numbers ovewflow errors are not covered

## Limitations
This library is for general purpose, but I want to maintain it compatible with https://github.com/lorenzotinfena/competitive-go so all the code should respect the limitations explained in the library. Maybe in the future the limitations could become less strict.