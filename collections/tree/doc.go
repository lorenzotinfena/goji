/*
By Q I mean the a result query of a segment.
In every possible range update to update a whole segment with result Q, in some cases is sufficient the old Q or not.

Example:
7 2 3 4 5 6 1 2
When you query (1-indexed) [1, 5] you combine the Q in range [1, 4] and the Q in range [5,5] and for max element you 7, and for the gcd you get 1.
When you update make a range update for example adding a positive integer in case of the maximum query is sufficient to get the old two Q, but in case of the gcd you need to reconstruct the segments, and lazy propagation in this last case is not worth.

Assuming that a simple Query to an element is O(1)
Summaring there are 3 types of updates:
1) Classic updates
2) Range updates with lazy propagation.
if there exist a function func(Qold)->Qnew in o(length of subsegment). And this can be implemented with lazy propagation in O(logN)
3) Range updates without lazy propagation
otherwise you need the Q of the two 2 subsegments, but to get those you need to iterate through all subsegments, which leads to an total for whole segment of O(length of subsegment), this explains also the fact that if there exist a function func(Qold)->Qnew not in o(length of subsegment) is convenient to follow this type of update.
This range update has a cost of (logN+(length of subsegment))

The first and third types can be unified in the same data structure because both they can store the elements other than the queries.
This is implemented by SegmentTree

The second data structure doesn't need to store the elements.
This is implemented by LazySegmentTree
*/
package tree
