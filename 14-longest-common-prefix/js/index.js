/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function(strs) {
    let str = strs[0];

    let res = "";

    for (let i = 0; i < str.length; i++) {
        for (let j = 0; j < strs.length; j++) {
            if (str[i] !== strs[j][i]) {
                return res;
            }
        }
        res += str[i];
    }

    return res;
};

module.exports = {
    longestCommonPrefix,
};
