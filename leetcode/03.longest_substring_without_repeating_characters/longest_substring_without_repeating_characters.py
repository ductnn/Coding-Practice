class Solution:
    def lengthOfLongestSubstring(self, s):
        maxLength = 0
        temp = ""

        if s == "":
            return 0

        for i in s:
            if i not in temp:
                temp += i
                if len(temp) > maxLength:
                    maxLength = len(temp)
            else:
                temp = temp[temp.index(i)+1:] + i

        return maxLength


if __name__ == "__main__":
    print(Solution().lengthOfLongestSubstring("ductnnnnn"))
