var addBinary = function(a, b) {
    // Reverse strings
    const revA = a.split('').reverse();
    const revB = b.split('').reverse();
    let carry = 0;
    let result = '';

    // Iterate up to the largest string
    for (let i = 0; i < Math.max(a.length, b.length); i++) {
        // Add a zero in case the lenght of one of the numbers have exhausted
        const numA = revA[i] === undefined ? 0 : parseInt(revA[i]);
        const numB = revB[i] === undefined ? 0 : parseInt(revB[i]);
        
        // Calculate the result
        // sum bit by bit using mod (%)
        // 0 % 2 = 0
        // 1 % 2 = 1
        // 2 % 2 = 0
        // 3 % 2 = 1
        const total = numA + numB + carry;
        const res = total % 2;
        
        // Calculate carry
        carry = Math.floor(total / 2);
        
        // Add res to result
        result = res.toString() + result;
    }

    // If there is carry add it to the final result
    if (carry === 1) {
        result = carry.toString() + result;
    }
   
    return result;
}

console.log(addBinary('11', '1'));
console.log(addBinary('1010', '1011'));
console.log(addBinary(
    "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
    "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011",
));

// 100
// 10101
// 110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000
