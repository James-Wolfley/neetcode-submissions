class Solution:
    def maxArea(self, heights: List[int]) -> int:
        max_water = 0
        left = 0
        right = len(heights) - 1

        while left < right:
            vol = (right - left) * min(heights[left], heights[right])
            max_water = max(max_water, vol)

            if heights[left] < heights[right]:
                left += 1
            else:
                right -= 1

        return max_water