const arr = [1, 3, 7, 12, 21, 33, 34, 40, 43, 52, 67, 101];

// find 12

function binarySearch(nums, target) {
    let left = 0;
    let right = nums.length - 1;

    while (left < right) {
        let m = Math.floor((left + right)/2);

        if (nums[m] == target) {
            return m;
        } else if (nums[m] > target) {
            right = m;
        } else {
            left = m;
        }
    }

    return -1;
}

console.log(binarySearch(arr, 33)) // 5
console.log(binarySearch(arr, 43)) // 8
console.log(binarySearch(arr, 3)) // 1
