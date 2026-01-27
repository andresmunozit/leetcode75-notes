/**
 * @param {number[]} digits
 * @return {number[]}
 */
var plusOne = function(digits) {
    // iterate from left to right and keep track of the carry,
    // if carry equals zero at some point just return the rest of the array
    let carry = 0

    for (let i = digits.length - 1; i >= 0; i--) {
        let result = 0;

        if (i === digits.length - 1) {
            result = digits[i] + 1;
        } else {
            result = digits[i] + carry;
        }

        if (result > 9) {
            carry = 1;
            digits[i] = result - 10;
        } else {
            digits[i] = result;
            return digits;
        }

        if (i === 0) {
            if (carry === 0) {
                return digits;
            } else {
                const carryArr = [1];
                return carryArr.concat(digits);
            }
        }
    }
};

module.exports = {
    plusOne,
};
