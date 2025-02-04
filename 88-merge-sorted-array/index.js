const assert = require('assert');

/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function (nums1, m, nums2, n) {
  // nums1 = [a, b, c, 0, 0, 0]
  // nums2 = [d, e, f]
  // Three pointers:
  // r1 (reader) = right of the relevant part of the first array (3)
  // r2 (reader) = right of the second array (3)
  // w (writter) = it will write at the end of the first array (5)
  let r1 = m - 1;
  let r2 = n - 1;

  for (let w = m + n - 1; w >= 0; w--) {
    if (r1 >= 0 && r2 >= 0) {
      nums1[w] = nums1[r1] > nums2[r2] ? nums1[r1--] : nums2[r2--];
    } else if (r1 >= 0) {
      nums1[w] = nums1[r1--];
    } else if (r2 >= 0) {
      nums1[w] = nums2[r2--];
    }
  }
};

function test(tt) {
  tt.forEach(t => {
    merge(t.nums1, t.m, t.nums2, t.n);
    assert.deepStrictEqual(t.nums1, t.expected);
  });
}

const tt = [
  {
    nums1: [1, 2, 3, 0, 0, 0],
    m: 3,
    nums2: [2, 5, 6],
    n: 3,
    expected: [1, 2, 2, 3, 5, 6],
  },
  {
    nums1: [1],
    m: 1,
    nums2: [],
    n: 0,
    expected: [1],
  },
  {
    nums1: [0],
    m: 0,
    nums2: [1],
    n: 1,
    expected: [1],
  },
]

test(tt);
