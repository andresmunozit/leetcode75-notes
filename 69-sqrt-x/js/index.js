/**
 * @param {number} x
 * @return {number}
 */
var mySqrt = function(x) {
    if (x <= 1) return x;

    // Newton algorithm
    let guess = x / 2;
    let newGuess = (guess + x / guess) / 2;

    while (guess - newGuess > Math.abs(0.1)) {
        guess = newGuess;
        newGuess = (guess + x / guess) / 2;
    }

    return Math.floor(newGuess);
};


console.log(mySqrt(4))
console.log(mySqrt(8))
console.log(mySqrt(16))
console.log(mySqrt(64))

// 2
// 2
// 16
// 64
