# https://leetcode.com/problems/successful-pairs-of-spells-and-potions

from bisect import bisect_left
from typing import List


class Solution:
    def successfulPairs(self, spells: List[int], potions: List[int], success: int) -> List[int]:
        potions.sort()
        return [len(potions) - bisect_left(potions, (success + spell - 1) // spell) for spell in spells]

    def bisect_left(self, nums: List[int], target: int) -> int:
        left, right = 0, len(nums)
        while left < right:
            mid = (left + right) // 2
            if nums[mid] < target:
                left = mid + 1
            else:
                right = mid
        return left

if __name__ == "__main__":
    spells = [5, 1, 3]
    potions = [1, 2, 3, 4, 5]
    success = 7
    print(Solution().successfulPairs(spells, potions, success))
    print(Solution().bisect_left(potions, success))
