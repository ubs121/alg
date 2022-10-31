# https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/
from typing import List

class Solution(object):
    def minimumDeleteSum(self, s1, s2):
        n,m=len(s1),len(s2)

        # Let dp[i][j] be the answer to the problem for the strings s1[i:], s2[j:].
        dp = [[0] * (m + 1) for _ in range(n + 1)]

        # When one of the input strings is empty, the answer is the ASCII-sum of the other string
        for i in range(n-1, -1, -1):
            dp[i][m] = dp[i+1][m] + ord(s1[i])
        for j in range(m-1, -1, -1):
            dp[n][j] = dp[n][j+1] + ord(s2[j])

        for i in range(n-1, -1, -1):
            for j in range(m-1, -1, -1):
                if s1[i] == s2[j]: 
                    # we have dp[i][j] = dp[i+1][j+1] as we can ignore these two characters.
                    dp[i][j] = dp[i+1][j+1]
                else: 
                    # we will have to delete at least one of them. We'll have dp[i][j] as the minimum of the answers after both deletion options.
                    dp[i][j] = min(dp[i+1][j] + ord(s1[i]),
                                   dp[i][j+1] + ord(s2[j]))

        return dp[0][0] # answer

def test_solution():
    test_cases = {
        "1": ( "sea", "eat", 231),
        "2": ( "delete", "leet", 403),
        "3": ( "ulaanbaatar", "naran", 1034),
        "4": ( "nbaatar", "naran", 615),
    }
    sol = Solution()
    for tc, (s1,s2, exp) in test_cases.items():
        got = sol.minimumDeleteSum(s1,s2)
        if got != exp:
            print(f"Test {tc}: exp {exp}, got {got}")
        else:
            print(f"Test {tc}: passed.")

if __name__ == '__main__':
    test_solution()