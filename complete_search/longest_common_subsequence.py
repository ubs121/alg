# https://en.wikipedia.org/wiki/Longest_common_subsequence_problem
# subsequences are not required to occupy consecutive positions within the original sequences
def LCSLength(s1: str, s2: str) -> int:
    n,m=len(s1),len(s2)
    lcs = [[0] *(m+1) for _ in range(n+1)]

    for i in range(1,n+1):
        for j in range(1,m+1):
            if s1[i-1] == s2[j-1]:
                lcs[i][j]=lcs[i-1][j-1] + 1
            else:
                lcs[i][j]=max(lcs[i][j-1], lcs[i-1][j])
    return lcs[n][m]

def test_solution():
    test_cases = {
        "1":("HARRY","SALLY",       2),
		"2":("AA","BB",             0),
		"3":("SHINCHAN","NOHARAAA", 3),
		"4":("HHNAN","NHAAA",       2),
	    "5":("AAABBB","BBBAAA",     3),
	    "6":("CAAAAAB","BAAAAAC",   5),
    }

    for tc, (s1,s2, exp) in test_cases.items():
        got = LCSLength(s1,s2)
        if got != exp:
            print(f"Test {tc}: exp {exp}, got {got}")
        else:
            print(f"Test {tc}: passed.")

if __name__ == '__main__':
    test_solution()