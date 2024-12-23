from typing import List

class Solution:
    def minimumDifference(self, nums: List[int], k: int) -> int:
        if k == 1:
            return 0  # If k is 1, the difference is always 0
        
        # Sort the scores
        nums.sort()
        
        # Initialize the minimum difference
        min_diff = float('inf')
        
        # Slide over the sorted array to find the minimum difference
        for i in range(len(nums) - k + 1):
            diff = nums[i + k - 1] - nums[i]
            min_diff = min(min_diff, diff)
        
        return min_diff

if __name__ == "__main__":
    solution = Solution()
    nums = [9, 4, 1, 7]
    k = 2
    result = solution.minimumDifference(nums, k)
    print(result)
