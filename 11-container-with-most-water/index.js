/**
 * Determines the maximum water volume between pairs of lines in `height`.
 * Implements a two-pointer approach, moving from the edges to the center.
 * The water volume is calculated as the min height of the two lines times the width.
 * @param {number[]} height - Heights of the lines.
 * @return {number} - Maximum water volume that can be contained.
 */
var maxArea = function(height) {
    const n = height.length

    // We initialize the container limits with two pointers
    let left = 0 // Index of the left limit (pointer) of the container
    let right = n - 1 // Index of the right limit of the container
    let maxArea = 0

    while (left < right) {
        const base = right - left
        const area = Math.min(height[left], height[right]) * base

        maxArea = Math.max(area, maxArea)

        // This is the most critical part, which pointer to move?
        // The smaller height limits the max area of the container, so we shift the
        // pointer with the smaller height to the left or right, searching for a bigger
        // height
        if (height[left] < height[right]) {
            left++
        } else {
            right--
        }
    }

    return maxArea
};

module.exports = maxArea
