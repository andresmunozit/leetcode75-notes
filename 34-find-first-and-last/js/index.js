function searchRange(nums, target) {
    // const left = binSearchLeft(nums, target);
    const left = binSearchLeft(nums, target);
    const right = binSearchRight(nums, target);

    return [left, right];
}

/**
 * @param {number[]} nums - The sorted array to search in
 * @param {number} target - The target value to find
 * @param {boolean} isLeft - Whether to find the leftmost occurrence
 */
function binSearchLeft(nums, target, isLeft) {
    let l = 0;
    let r = nums.length - 1;
    let i = -1;
    
    // l and r could converge
    while (l <= r) {
        const m = Math.floor((l + r) / 2);
        if (nums[m] < target) {
            l = m + 1;
        } else {
            r = m - 1;
        }
        if (nums[m] == target) {
            i = m;
        }
    }

    return i
}

function binSearchRight(nums, target) {
    let l = 0;
    let r = nums.length - 1;
    let i = -1;
    
    while (l <= r) {
        const m = Math.floor((l + r + 1)/2);
        if (nums[m] > target) {
            r = m - 1;
        } else {
            l = m + 1;
        }
        if (nums[m] == target) {
            i = m;
        }
    }

    return i
}

console.log(searchRange([5,7,7,8,8,10], 8)); // [3, 4]
console.log(searchRange([5,7,7,8,8,10], 6)); // [-1, -1]
console.log(searchRange([], 0)); // [-1, -1]
