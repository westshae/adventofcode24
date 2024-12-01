1. Based on location id
2. Two groups go around creating their own complete list
3. The two lists aren't the same.
4. Need to load data from a .txt, with {location id 1}{3 whitespace}{location id 2}
5. For each pair of ids, find the diff between them (3 & 6 == 3, 7 & 3 == 4)
6. For each side, you must sort it in order of smallest to largest, then find the diff. ({3,1,2} && {6,2,10} === {1,2,3} && {2, 6, 10} === {1, 4, 7} === 12)
6. Find the total distance, as shown above.