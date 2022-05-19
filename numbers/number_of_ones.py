# https://leetcode.com/problems/number-of-digit-one/
# Given an integer n, count the total number of digit 1 appearing in all non-negative integers less than or equal to n.

# F - number of ones in "d digit" number
# F(d) = Q(d)-Q(d-1)
# Q(d) = sum(k*C(d,k)*pow(9,d-k)), where k=[1,d], C(d,k) - combinations of k from d digits