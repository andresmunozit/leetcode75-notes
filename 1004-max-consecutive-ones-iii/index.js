/**
 * Finds the maximum number of consecutive 1s in the array if you can flip at most k 0s.
 * @param {number[]} nums - Binary array (only contains 0s and 1s).
 * @param {number} k - The maximum number of 0s that can be flipped to 1s.
 * @return {number} - The maximum number of consecutive 1s after flipping at most k 0s.
 */
var longestOnes = function (nums, k) {
  let left = 0; // Left pointer for the sliding window
  let right = 0; // Right pointer for the sliding window
  let longest = 0; // Track the length of the longest subarray
  let zerosCount = 0; // Count the number of zeros in the current window

  // Iterate through the array with the right pointer
  while (right < nums.length) {
    // Increment zerosCount if current element is 0
    if (nums[right] === 0) {
      zerosCount++;
    }

    // When zerosCount exceeds k, update longest and shrink window from the left
    if (zerosCount > k) longest = Math.max(longest, right - left);
    while (zerosCount > k) {
      // If the left element is 0, decrement zerosCount
      if (nums[left] === 0) {
        zerosCount--;
      }
      // Move the left pointer to shrink the window
      left++;
    }

    // Move the right pointer to expand the window
    right++;
  }

  // Update longest for the case where the loop exits without updating longest
  longest = Math.max(longest, right - left);

  return longest;
};

var a = longestOnes([1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0], 2) // expected 6

module.exports = longestOnes
