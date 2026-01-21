var isPalindrome = function(x) {
    // check if x < 0 or x % 10 = 0, in such cases x can't be a palindrom
    if (x < 0 || (x % 10 === 0 && x !== 0)) {
        return false;
    }

    let reversed = 0;

    // start eliminating the less significative digit and add it to a reversed while original > reversed
    while (x > reversed) {
        reversed = reversed * 10 + x % 10;
        x = Math.floor(x / 10);
    }

    // check if x == reversed or x == reversed // 10
    return x === reversed || x === Math.floor(reversed / 10);
};

console.log(isPalindrome(121))
console.log(isPalindrome(-121))
console.log(isPalindrome(10))
// true
// false
// false
