/**
 * @param {number[]} arr
 * @return {boolean}
 */
var uniqueOccurrences = function(arr) {
    const numMap = {}
    const freqMap = {}

    for (let n of arr) {
        numMap.hasOwnProperty(n) ? numMap[n]++ : numMap[n] = 1
    }

    for (let k in numMap) {
        const freq = numMap[k]
        freqMap.hasOwnProperty(freq) ? freqMap[freq]++ : freqMap[freq] = 1
        if (freqMap[freq] > 1) {
            return false
        }
    }

    return true
};

var output = uniqueOccurrences([1,2])

module.exports = uniqueOccurrences
