# https://leetcode.com/problems/edit-distance/

# Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.

# You have the following three operations permitted on a word:

# Insert a character
# Delete a character
# Replace a character

class Solution:
    def min_distance(self, word1: str, word2: str) -> int:
        memo={}
        def edit_distance(i: int, j: int):
            if i<0 or j<0:
                return i+1+j+1
                
            if (i,j) in memo:
                return memo[(i,j)]
            
            if word1[i]==word2[j]:
                result=edit_distance(i-1,j-1)
            else:
                result=1+min(
                            edit_distance(i-1,j),  # delete
                            edit_distance(i, j-1), # insert
                            edit_distance(i-1,j-1) # replace
                            )
            memo[(i,j)]=result
            return result

        return edit_distance(len(word1)-1,len(word2)-1)

def test_solution():
    test_cases = {
        "1": ( "horse", "ros", 3),
        "2": ( "intention", "execution", 5),
        "3": ( "", "naran", 5),
    }
    sol = Solution()
    for tc, (s1,s2, exp) in test_cases.items():
        got = sol.min_distance(s1,s2)
        if got != exp:
            print(f"Test {tc}: exp {exp}, got {got}")
        else:
            print(f"Test {tc}: passed.")

if __name__ == '__main__':
    test_solution()