/**
 * Finds the maximum number of k-sum pairs in the array.
 * @param {number[]} nums - The array of numbers.
 * @param {number} k - The target sum for the pairs.
 * @return {number} - The maximum number of k-sum pairs.
 */
var maxOperations = function(nums, k) {
    const req = {}; // Object to store the required numbers to form a pair with each element.
    let maxCount = 0; // Counter for the maximum number of k-sum pairs.
    
    for (const n of nums) {
        const complement = k - n; // Calculate the complement number needed to reach k.

        // If the complement exists in req, a pair is formed.
        if (req[complement] >= 1) {
            req[complement]--; // Decrement the count of the complement as it's now used.
            maxCount++; // Increment the count of pairs.
        } else {
            // If the complement is not found, record the current number.
            req[n] = (req[n] || 0) + 1;
        }
    }

    return maxCount; // Return the total number of pairs formed.
};

// Example usage
// var m = maxOperations([3,1,3,4,3], 6);

module.exports = maxOperations
