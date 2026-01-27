/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var searchInsert = function(nums, target) {
    if (nums.length === 0) {
        return 0;
    }

    if (target < nums[0]) {
        return 0;
    }

    if (target > nums[nums.length - 1]) {
        return nums.length;
    }

    let result;

    for (let i = 0; i < nums.length; i++) {
        if (nums[i] === target) {
            result = i;
            break;
        } else if (i > 0 && target > nums[i-1] && target < nums[i]) {
            result = i;
            break;
        }
    }

    return result;
};

module.exports = {
    searchInsert,
};
