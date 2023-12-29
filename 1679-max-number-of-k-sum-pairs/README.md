# 1679. Max Number of K-Sum Pairs  
Given an integer array `nums` and an integer `k`, the task is to find the maximum number of
operations where each operation consists of removing a pair of numbers from the array that add up to
`k`.  

## Examples  

### Example 1:  
Input: `nums` = [1,2,3,4], `k` = 5

Output: 2 

Explanation:   
- Remove numbers 1 and 4, making `nums` = [2,3].
- Remove numbers 2 and 3.

After this, there are no more pairs that sum up to 5, resulting in a total of 2 operations.  

### Example 2:  
Input: `nums` = [3,1,3,4,3], `k` = 6

Output: 1

Explanation:  
- Remove the first two 3's, making `nums` = [1,4,3].

There are no more pairs that sum up to 6, hence a total of 1 operation.

## Constraints  
- 1 <= `nums.length` <= 10^5  
- 1 <= `nums[i]` <= 10^9  
- 1 <= `k` <= 10^9  
