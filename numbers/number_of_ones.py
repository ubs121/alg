# https://leetcode.com/problems/number-of-digit-one/
# Given an integer n, count the total number of digit 1 appearing in all non-negative integers less than or equal to n.
import math

class Solution:
    def countDigitOne(self, n: int) -> int:
        if n==0: return 0

        n_str=str(n)
        total_1s=0
        prefix_1s=0
        for i, c in enumerate(n_str):
            d=int(c)
            p=len(n_str)-i-1
            curr=d*pow(10, p)
            
            # accumulate total 1s
            total_1s+=prefix_1s*curr+self.ones(d, p)

            if d==1:
                prefix_1s+=1

        return total_1s
    
    # returns all 1s in 'p' placeholders
    # input: d - first digit, p - number of zeros after 'd'
    def ones(self, d: int, p: int) -> int:
        if d==0:
            return 0
        if p==0:
            return 1  # d>0

        # 1s in up to d*10^n numbers = sum(k*C(n,k)*pow(9,n-k)), where k=[1,n], C(n,k) - combinations
        s=sum([k*math.comb(p,k)*pow(9,p-k) for k in range(1,p+1)])
        if d==1:
            return s+1
        return d*s+pow(10, p)
    
    def countDigitOneBF(self, n: int) -> int:
        m=0
        for i in range(n+1):
            str_num=str(i)
            for c in str_num:
                if c=='1':
                    m+=1
        return m


def test_solution():
    test_cases = {
        "0": (0, 0),
        "1": (1, 1),
        "13": (13, 6),
        "10": (10, 2),
        "30": (30, 13),
        "100": (100, 21),
        "200":(200,140),
        "1000": (1000, 301),
        "300": (300, 160),
        "1979": (1979, 1578),
        "2000": (2000,1600),
        "1000000000":(1000000000,900000001),
    }
    sol = Solution()
    for tc, (n, exp) in test_cases.items():
        got = sol.countDigitOne(n)
        if got != exp:
            print(f"Test {tc}: exp {exp}, got {got}")
        else:
            print(f"Test {tc}: passed.")

if __name__ == '__main__':
    test_solution()