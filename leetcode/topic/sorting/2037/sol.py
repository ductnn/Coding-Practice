from typing import List

class Solution:
    def minMovesToSeat(self, seats: List[int], students: List[int]) -> int:
        seats.sort()
        students.sort()

        total = 0
        for i in range(len(seats)):
            total += abs(seats[i] - students[i])
        
        return total
    
if __name__ == "__main__":
    solution = Solution()
    seats = [4, 1, 5]
    students = [2, 7, 3]
    print(solution.minMovesToSeat(seats, students))
