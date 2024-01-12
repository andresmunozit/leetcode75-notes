/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var findMaxAverage = function(nums, k) {
    // Sliding window approach
    let sum = 0
    let maxAv = 0

    // Calculate the initial sum
    for (let i = 0; i < k; i++) {
        sum += nums[i]
    }

    maxAv = sum / k

    // Now let's slide the window
    for (let i = k; i < nums.length; i++) {
        sum = sum - nums[i-k] + nums[i]
        const a = sum / k
        if (a > maxAv) maxAv = a
    }

    return maxAv
};

// Use example
var av = findMaxAverage([1,12,-5,-6,50,3], 4)

module.exports = findMaxAverage
