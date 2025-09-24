# https://leetcode.com/problems/fraction-to-recurring-decimal/

class Solution:
    def fractionToDecimal(self, numerator: int, denominator: int) -> str:
        sign = ""
        if numerator < 0 and denominator > 0 or numerator > 0 and denominator < 0:
            sign = "-"
        numerator = abs(numerator)
        denominator = abs(denominator)

        integerPart = numerator // denominator
        print("integerPart", integerPart)
        remainder = numerator % denominator
        print("remainder", remainder)

        fractionPart = ""
        if remainder != 0:
            fractionPart = "."
        
        remainderMap = {}
        repeatIndex = -1

        while remainder != 0:
            if remainder in remainderMap:
                repeatIndex = remainderMap[remainder]
                print("repeatIndex", repeatIndex)
                break
            remainderMap[remainder] = len(fractionPart)
            remainder *= 10
            fractionPart += str(remainder // denominator)
            remainder %= denominator
            print("fractionPart", fractionPart)
        if repeatIndex != -1:
            fractionPart = fractionPart[:repeatIndex] + "(" + fractionPart[repeatIndex:] + ")"
            print("fractionPart-after", fractionPart)

        return sign + str(integerPart) + fractionPart

    def abs(self, a: int) -> int:
        if a < 0:
            return -a
        return a

if __name__ == "__main__":
    solution = Solution()
    print(solution.fractionToDecimal(-1, 2))
    print(solution.fractionToDecimal(2, 1))
    print(solution.fractionToDecimal(4, 333))
