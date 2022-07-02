# https://leetcode.com/problems/word-search/
from collections import defaultdict, deque
from typing import List

class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        self.n=len(board)
        self.m=len(board[0])
        self.board=board
        self.word=word
        self.failures=defaultdict(int)
        
        # check from each cell
        for r in range(self.n):
            for c in range(self.m):
                #print("Checking at", (r,c))
                self.failures.clear() # reset failure memory
                if board[r][c]==word[0]:
                    (complete,_)=self.checkAt(r,c,0)
                    if complete:
                        return True
        
        # doesn't exist if search is done
        return False

    def checkAt(self, r, c, p):
        if p<len(self.word) and (0<=r<self.n and 0<=c<self.m) and self.word[p]==self.board[r][c] and self.board[r][c]!="*":
            letter=self.word[p]

            # check from the failures
            if self.failures[(r,c,p)]>=pow(3, len(self.word)-p): # this is problematic ! how many failures are acceptable ?
                #print("was failed at", (r,c,self.word[p:]), self.failures[(r,c,p)], " times")
                return (False,False)

            #print(letter)
            self.board[r][c]="*" # mark it visited

            # check neighbors for next letter
            for d in [(1,0),(0,1),(-1,0),(0,-1)]: 
                (complete,valid)=self.checkAt(r+d[0], c+d[1], p+1)
                if complete:
                    return (complete,valid) # word is complete

            # reverse
            #print("["+letter+"]") 
            self.board[r][c]=letter

            # learn the failure at (r,c,p)
            if p<len(self.word)-1:
                #print("failed at", (r,c,self.word[p:]))
                self.failures[(r,c,p)]+=1
            return (p==len(self.word)-1, True)
        return (False,False)

    def checkBSF(self, r, c, p):
        source=(r,c)
        parent, Q = {source:None}, deque([(source,p)])
        while Q:
            (u,w)=Q.popleft() # dequeue
            #print(self.board[u[0]][u[1]],",",self.word[w:])
            if w==len(self.word)-1:
                return (True, True)

            for v in [(u[0]+1, u[1]),(u[0]-1, u[1]),(u[0], u[1]+1),(u[0], u[1]-1)]: # all neighbors
                if not (0<=v[0]<self.n and 0<=v[1]<self.m):  continue # invalid boundary
                if self.board[v[0]][v[1]]!=self.word[w+1]:  continue  # unexpected letter

                if v in parent: 
                    continue # already visited

                parent[v]=u # mark it seen
                Q.append((v, w+1))
        return (False,False)
        
def test_solution():
    testCases={
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
    for tc,(board,word,exp) in testCases.items():
        got=sol.exist(board, word)
        if got!=exp:
            print("Test {}: exp {}, got {}".format(tc, exp, got))
        else:
            print("Test {}: passed.".format(tc))

if __name__ == '__main__':
    test_solution()