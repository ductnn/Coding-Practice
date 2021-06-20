######################
## Author: Duc Tran ##
## Title: Two Sum   ##
######################

class Solution:
    def twoSum(self, nums, target):
        for i in range(len(nums)):
            for j in range(i+1, len(nums)):
                if nums[i] + nums[j] == target:
                    return [i, j]

    def twoSum1(self, nums, target):
        values = {}
        for i, num in enumerate(nums):
            if (target - num) in values:
                return [values[target - num], i]
            values[num] = i
        return []

if __name__ == '__main__':
    solution = Solution()
    nums = list([2, 7, 11, 15])
    target = 9
    print(solution.twoSum1(nums, target))
