# https://leetcode.com/problems/word-search/
from collections import defaultdict
from typing import List, Tuple

class Solution:
    def __init__(self) -> None:
        self.failures=defaultdict(int)

    def exist(self, board: List[List[str]], word: str) -> bool:
        self.n=len(board)
        self.m=len(board[0])
        self.board=board
        self.word=word
        self.failures.clear() # reset failure memory

        # check from each cell
        for r in range(self.n):
            for c in range(self.m):
                #print("Checking at", (r,c))
                
                (complete,_)=self.check_dfs((r,c),0)
                if complete:
                    return True
        
        # doesn't exist if search is done
        return False

    # DFS search starting from source
    def check_dfs(self, u: Tuple, p: int) -> Tuple:
        (x,y)=u
        if p<len(self.word) and (0<=x<self.n and 0<=y<self.m) and self.word[p]==self.board[x][y] and self.board[x][y]!="*":
            letter=self.word[p]

            # Prune: check from the failures
            # going forward each remaining cells have 3 options, so 3^r (r - # of remaining letters) options should be checked to conclude it's a failure
            if self.failures[(u,p)]>=pow(3, len(self.word)-p):
                #print("was failed at", (r,c,self.word[p:]), self.failures[(r,c,p)], " times")
                return (False,False)

            #print(letter)
            self.board[x][y]="*" # mark it visited

            # check neighbors for next letter
            for v in [(x+1, y),(x-1, y),(x, y+1),(x, y-1)]:
                (complete,valid)=self.check_dfs(v, p+1)
                if complete:
                    return (complete,valid) # word is complete

            # learn the failure at (r,c,p)
            if p<len(self.word)-1:
                #print("failed at", (r,c,self.word[p:]))
                self.failures[(u,p)]+=1
                # reverse the visited '*' mark
                #print("["+letter+"]")
                self.board[x][y]=letter
            return (p==len(self.word)-1, True)
        return (False,False)

def test_solution():
    test_cases={
        "1":([["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], "ABCCED", True),
        "2":([["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], "SEE", True),
        "3":([["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], "ABCB", False),
        "4":([["a","a"]], "aaa", False),
        "5":([["a","b"],["c","d"]], "acdb", True),
        "6":([["a","b"],["c","d"]], "abcd", False),
        "7":([["A","B","C","E"],["S","F","E","S"],["A","D","E","E"]], "ABCESEEEFS", True),
        "8":([["a"]], "a", True),
        "9":([["C","A","A"],["A","A","A"],["B","C","D"]], "AAB", True),
        "10":([["A","A","A","A","A","A"],["A","A","A","A","A","A"],["A","A","A","A","A","A"],["A","A","A","A","A","A"],["A","A","A","A","A","B"],["A","A","A","A","B","A"]], "AAAAAAAAAAAAABB", False),
        "11":([["A","Z","A","A"],["A","B","B","B"],["A","C","B","B"],["A","B","B","B"],["A","B","B","B"]],"BBBBBBBCBZ",True),
        "12":([["A","A","a","a","A","a"],["a","a","a","A","A","a"],["A","a","A","a","a","A"]],"AAaaAAaAaaAaAaA",True),
    }
    sol=Solution()
    for tc,(board,word,exp) in test_cases.items():
        got=sol.exist(board, word)
        if got!=exp:
            print(f"Test {tc}: exp {exp}, got {got}")
        else:
            print(f"Test {tc}: passed.")

if __name__ == '__main__':
    test_solution()
