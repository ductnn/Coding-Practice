class Solution:
    def findMinDifference(self, timePoints: list[str]) -> int:
        # Convert time to minutes since "00:00"
        def timeToMinutes(time: str) -> int:
            hours, minutes = map(int, time.split(":"))
            return hours * 60 + minutes
        
        # Convert all time points to minutes and sort them
        minutes_list = sorted(timeToMinutes(time) for time in timePoints)
        
        # Initialize minimum difference with a large number
        min_diff = float('inf')

        # Compare consecutive times
        for i in range(1, len(minutes_list)):
            min_diff = min(min_diff, minutes_list[i] - minutes_list[i - 1])

        # Compare the first and last time, considering the circular nature of time
        min_diff = min(min_diff, (24 * 60 - minutes_list[-1] + minutes_list[0]))

        return min_diff

if __name__ == "__main__":
    solution = Solution()
    timePoints = ["23:59", "00:00"]
    result = solution.findMinDifference(timePoints)
    print(result)
