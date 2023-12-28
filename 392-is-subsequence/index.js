/**
 * Determines if `s` is a subsequence of `t`. A subsequence is formed by deleting some or no 
 * elements from a string without changing the order of the remaining elements. This function 
 * iterates through `t` and compares it with `s`, increasing a pointer for each match. 
 * If all characters of `s` are found in sequence in `t`, `s` is a subsequence of `t`.
 * 
 * @param {string} s - String to check as a subsequence.
 * @param {string} t - String to be checked against.
 * @return {boolean} - True if `s` is a subsequence of `t`, false otherwise.
 */

var isSubsequence = function(s, t) {
    // Traverse t and advance i when s[i] is found
    let i = 0 // Index for iterate over t
    let j = 0 // Index for iterate over s

    // Edge case
    if (s.length > t.length) {
        return false
    }

    while (i < t.length && j < s.length) {
        if (s[j] === t[i]) {
            j++
        }
        i++
    }

    return j === s.length
};

module.exports = isSubsequence
