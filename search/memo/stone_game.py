# https://leetcode.com/problems/stone-game/
from collections import defaultdict
from typing import List

# even cuts - Alice, odd cuts - Bob
# [l:r] - memorize the winner for each range
class Solution:
    def stoneGame(self, piles: List[int]) -> bool:
        return True # Alice can manage to win always, because # number of piles are even

    def stoneGame2(self, piles: List[int]) -> bool:
        self.mem=defaultdict(bool)
        self.piles=piles
        self.total=sum(piles)
        ans=self.solve(0, len(piles)-1, 0, 0, 0)
        print("mem=", len(self.mem), "len=", len(piles))
        return  ans
        
    # a- alice stones, b-bob stones, t-turn
    def solve(self, l:int, r:int, a:int, b:int, t: int) -> bool:
        if l>r:
            return a>b
        
        if (l,r,a-b) in self.mem:
            return self.mem[(l,r,a-b)]

        if a>self.total/2:
            return True
        if b>self.total/2:
            return False

        # Alice
        if t%2==0:
            beg=self.solve(l+1,r, a+self.piles[l], b, t+1)
            if beg:
                self.mem[(l,r,a-b)]=True
                return True
            end=self.solve(l,r-1, a+self.piles[r], b, t+1)
            self.mem[(l,r,a-b)]=(beg or end)
            return beg or end # Alice can win if one of choices are True
        
        # Bob
        beg=self.solve(l+1,r, a, b+self.piles[l], t+1)
        if not beg:
            self.mem[(l,r,a-b)]=False
            return False
        end=self.solve(l,r-1,a, b+self.piles[r], t+1)
        self.mem[(l,r,a-b)]=(beg and end)
        return beg and end # Bob can win if one of choices are False

    # reverse direction
    def solve2(self, l:int, r:int, a:int, b:int, t: int) -> bool:
        # start from ending and accumulate (l,r) and (a,b)
        # it's easy because each player would choose max from two ends to maximize the score
        
        t=1 # Bob ends
        b+=self.piles[l]
        l-=1
        while t<len(self.piles) and 0<=l and r<len(self.piles):
            amt=0
            if self.piles[l]<self.piles[r]:
                amt=self.piles[r]
                r+=1
            else:
                amt=self.piles[l]
                l-=1

            # accumulate to player's score
            if t%2==0:
                a+=amt
            else:
                b+=amt
            
            t+=1

        if t<len(self.piles):
            if l<0:
                # right
                while r<len(self.piles):
                    if t%2==0:
                        a+=self.piles[r]
                    else:
                        b+=self.piles[r]
                    t+=1
                    r+=1
            else: # left
                while l>=0:
                    if t%2==0:
                        a+=self.piles[l]
                    else:
                        b+=self.piles[l]
                    t+=1
                    l-=1
        print(f"a={a}, b={b}")
        return a>b

def test_solution():
    test_cases = {
        #"0": ([5,3,4,5], True),
        #"1": ([10,3,1,7,8,6,8,10], True),
        #"2": ([3,7,2,3], True),
        "3":([34,100,35,29,95,3,76,26,81,48,61,4,30,90,31,21,16,70,40,46,30,76,40,25,92,99,10,12,70,82,62,98,14,68,94,5,9,64,34,89,98,54,41,56,60,30,4,38,67,76,71,40,89,83,19,49,97,97,61,95,6,55,14,34,35,44,68,51,32,93,36,98,87,79,29,46,46,8,75,18,63,9,52,60,3,76,89,86,4,22,7,30,93,31,52,28,51,74,95,60],True),
    }
    sol = Solution()
    for tc, (arr, exp) in test_cases.items():
        got = sol.stoneGame(arr)
        if got != exp:
            print(f"Test {tc}: exp {exp}, got {got}")
        else:
            print(f"Test {tc}: passed.")

if __name__ == '__main__':
    test_solution()
    # sol=Solution()
    # piles=[10,3,1,7,8,6,8,10]
    # sol.stoneGame(piles)
    # #for i in range(len(piles)):
    # i=2
    # ans=sol.solve2(i,i, 0, 0, 1)
    # print(f"i={i}", ans)
