var addBinary = function(a, b) {
    return (parseInt(a, 2) + parseInt(b, 2)).toString(2);
}

console.log(addBinary('11', '1'));      // 100
console.log(addBinary('1010', '1011')); // 10101
