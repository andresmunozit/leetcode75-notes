/**
 * @param {character[]} chars
 * @return {number}
 */
/**
 * @param {character[]} chars
 * @return {number}
 */
var compress = function(chars) {
    // Intial conditions
    const compressed = [chars[0]]
    let prevChar = chars[0]
    let charCount = 1

    // Traverse the array, if the current char is the same previous one, add a counter
    for (let i = 1; i <= chars.length; i++) {
        if (chars[i] !== prevChar) {
            // If `charCount` is multicharacter, split it in (i.e. '12' => '1', '2')
            if (charCount.toString().length > 1) {
                compressed.pop()
                compressed.push(...Array.from(charCount.toString()))
            }

            if (i === chars.length) {
                break
            }

            prevChar = chars[i]
            charCount = 1
            compressed.push(chars[i])
            continue
        }

        charCount += 1
        if (charCount === 2) {
            compressed.push(charCount.toString())
            continue
        }
       
        compressed[compressed.length - 1] = charCount.toString()
    }
    
    // Adds the compressed array at the begginning of the chars array
    chars.unshift(...compressed)
    return compressed.length
};

module.exports = compress
