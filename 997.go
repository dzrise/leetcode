/*
997. Find the Town Judge
Medium
Topics
Companies
In a town, there are n people labeled from 1 to n. There is a rumor that one of these people is secretly the town judge.

If the town judge exists, then:

The town judge trusts nobody.
Everybody (except for the town judge) trusts the town judge.
There is exactly one person that satisfies properties 1 and 2.
You are given an array trust where trust[i] = [ai, bi] representing that the person labeled ai trusts the person labeled bi. If a trust relationship does not exist in trust array, then such a trust relationship does not exist.

Return the label of the town judge if the town judge exists and can be identified, or return -1 otherwise.

Example 1:

Input: n = 2, trust = [[1,2]]
Output: 2
Example 2:

Input: n = 3, trust = [[1,3],[2,3]]
Output: 3
Example 3:

Input: n = 3, trust = [[1,3],[2,3],[3,1]]
Output: -1

Constraints:

1 <= n <= 1000
0 <= trust.length <= 104
trust[i].length == 2
All the pairs of trust are unique.
ai != bi
1 <= ai, bi <= n
*/
package main

import "fmt"

func findJudge(n int, trust [][]int) int {
	if len(trust) == 0 {
		if n == 1 {
			return 1
		}
		return -1
	}
	var mapDov = make(map[int]int, n)
	var mapNoDov = make(map[int]bool, n)

	for i := 0; i < n; i++ {
		mapNoDov[i] = false
		mapDov[i] = 0
	}

	for i := 0; i < len(trust); i++ {
		mapNoDov[trust[i][0]] = true
		mapDov[trust[i][1]]++
	}

	max := 0
	res := 0
	for i, v := range mapDov {
		if v > max {
			max = v
			res = i
		}
	}

	if max == n-1 && !mapNoDov[res] {
		return res
	}

	return -1
}

func main() {
	fmt.Println([][]int{[]int{1, 2}})
	fmt.Println(findJudge(2, [][]int{[]int{1, 2}}))
}
