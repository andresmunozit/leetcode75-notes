/**
 * @param {string} s
 * @return {number}
 */
var romanToInt = function(s) {
    // create a hash map of values
    const values = {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    };

    sum = 0;

    // iterate over s, if curr < next subtract, otherwise add
    for (let i = 0; i < s.length; i++) {
        let curr = values[s[i]];

        if (i < s.length - 1 && curr < values[s[i+1]]) {
            sum -= curr;
        } else {
            sum += curr;
        }
    }

    return sum;
};

console.log(romanToInt('III'));
console.log(romanToInt('LVIII'));
console.log(romanToInt('MCMXCIV'));
// 3
// 58
// 1994
