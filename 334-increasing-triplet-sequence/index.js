/**
 * 
 * @param {*} nums 
 * @returns 
 */
var increasingTriplet = function(nums) {
    let i = null // The first smaller number
    let j = null // The second smaller number

    for (const num of nums) {
        // Look for the first smaller `num`
        if (i === null || num <= i) {
            // If we found it assign it to `i`
            i = num
            continue
        }

        // At this point is waranteed that i exists and the current `num`
        // is greater than `i`
        if (j === null || num <= j) {
            j = num
            continue
        }

        // At thhis point is waranteed that `i` and `j` exist and
        // num > j > i, so we had found a triplet
        return true
    }

    return false
}

module.exports = increasingTriplet
