const arr = [1, 3, 7, 12, 21, 33, 34, 40, 43, 52, 67, 101];

function binarySearch(nums, target) {
    let left = 0;
    let right = nums.length - 1;

    // There can be the case in which left == rigth, win which case both will contain the index
    // of target
    while (left <= right) {
        // Integer division
        let m = Math.floor((left + right)/2);

        if (nums[m] == target) {
            return m;
        } else if (nums[m] > target) {
            // This excludes the most right element which we already compared
            right = m - 1;
        } else {
            // This excludes the most left element which we already compared
            left = m + 1;
        }
    }

    return -1;
}

console.log(binarySearch(arr, 3)) // 1
console.log(binarySearch(arr, 33)) // 5
console.log(binarySearch(arr, 43)) // 8
console.log(binarySearch(arr, 22)) // -1
console.log(binarySearch(arr, -3)) // -1
console.log(binarySearch(arr, 200)) // -1
