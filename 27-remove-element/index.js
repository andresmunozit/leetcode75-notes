const assert = require('assert');

/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
var removeElement = function (nums, val) {
  // Two-pointer approach: one pointer (w) writes, the other (r) reads.
  // The writer pointer (w) tracks the new valid array length.
  let w = 0;
  let r = 0;

  // Iterate through the array using the reader pointer (r).
  while (r < nums.length) {
    if (nums[r] === val) {
      // If the current element is the target value, skip it.
      r++;
    } else {
      // Copy the element to the writer's position and move both pointers.
      // The writer pointer (w) only moves when we write a non-target value.
      nums[w++] = nums[r++];
    }
  }

  // The final value of 'w' is the number of elements that are not 'val'.
  return w;
};

// Example usage

/*
  Time Complexity: O(n)
  - We traverse the array once with the reader pointer (r), so the time complexity is O(n).

  Space Complexity: O(1)
  - We modify the input array in-place without using extra space.
  - Only a few integer variables are used (constant space).
*/
