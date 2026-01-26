/**
 * @param {string} haystack
 * @param {string} needle
 * @return {number}
 */
var strStr = function(haystack, needle) {
    for (let i = 0; i <= haystack.length - needle.length; i++) {
        if (haystack.slice(i, needle.length + i) === needle) {
            return i;
        }
    }
    return -1;
};

module.exports = {
    strStr,
};
