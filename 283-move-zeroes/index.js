/**
 * Shifts all zeroes in `nums` to the end while keeping other elements in order.
 * Uses `left` and `right` pointers to swap zeroes with non-zero elements.
 * Example: [1,4,0,0,5] -> left = 2, right = 4. Swap nums[4] with nums[2].
 * Result: [1,4,5,0,0]
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var moveZeroes = function(nums) {
    let i = 0 // Will traverse the nums array
    let left = null // left limit of zeroes window
    let right = null // right limit of zeroes window

    const n = nums.length

    while (i < n && right < n) {
        if (nums[i] === 0) {
            // Initialize left and right when zero is found for the first time
            if (left === null) {
                left = i
                if (right === null) {
                    right = left + 1
                }
                i++
                continue
            }
            i++
            right++
            continue
        }

        // Continue looking for a zero
        if (left === null && right === null) {
            i++
            continue
        }

        // Swap num with the zero at the left
        [nums[i], nums[left]] = [nums[left], nums[i]]
        
        left++
        right++
        i = right
    }
};

module.exports = moveZeroes
