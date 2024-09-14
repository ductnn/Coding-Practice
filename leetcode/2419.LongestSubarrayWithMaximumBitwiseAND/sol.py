def longestSubarray(nums):
    max_val = max(nums)

    longest = 0
    current_length = 0

    for num in nums:
        if num == max_val:
            current_length += 1
            longest = max(longest, current_length)
        else:
            current_length = 0

    return longest

if __name__ == "__main__":
    nums = [1, 2, 3, 3, 2, 2, 3]

    result = longestSubarray(nums)
    print(result)
