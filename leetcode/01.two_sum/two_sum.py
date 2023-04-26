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

    # Sorted test
    def twoSum2(self, nums, target):
        i = 0
        j = len(nums) - 1

        while i < j:
            if nums[i] + nums[j] == target:
                return [i, j]
            elif nums[i] + nums[j] > target:
                j -= 1
            else:
                i +=1
        return []

if __name__ == '__main__':
    solution = Solution()
    nums = list([2, 7, 11, 15])
    target = 9
    print(solution.twoSum2(nums, target))
