const assert = require('assert')
const { benchmark } = require('../../utils/index')
/**
 * @param {string} s
 * @return {string}
 */
// This function will traverse the string from the left until a vowel appears
// Then swap with the vowel at the right
var reverseVowels = function(s) {
    let i = 0
    let j = s.length - 1
    const reverse = Array(s.length).fill()
    while (i <= j){
        if (!isVowel(s[i])) {
            reverse[i] = s[i]
            i++
            continue
        }
        while (j >= i) {
            if (!isVowel(s[j])) {
                reverse[j] = s[j]
                j--
                continue
            }

            reverse[i] = s[j]
            reverse[j] = s[i]
            j--
            // break this  while loop to continue from the left
            break
        }
        i++
    }

    return reverse.join('')
}

var reverseVowels2 = function(s) {
    const vowels = Array.from(s).filter(isVowel)
    const rev = []
    for (let i = 0; i < s.length; i++) {
        if (!isVowel(s[i])) {
            rev[i] = s[i]
            continue
        }
        rev[i] = vowels[vowels.length - 1]
        vowels.pop()
    }
    return rev.join('')
}

function isVowel(a) {
    return ['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'].includes(a)
} 

// Test `reverseVowels`
assert.strictEqual(reverseVowels('hello'), 'holle')
assert.strictEqual(reverseVowels('leetcode'), 'leotcede')
assert.strictEqual(reverseVowels('a.'), 'a.')
assert.strictEqual(reverseVowels('aA'), 'Aa')
benchmark(reverseVowels, 'leetcode')

// Test `reverseVowels2`
assert.strictEqual(reverseVowels2('hello'), 'holle')
assert.strictEqual(reverseVowels2('leetcode'), 'leotcede')
assert.strictEqual(reverseVowels2('a.'), 'a.')
assert.strictEqual(reverseVowels2('aA'), 'Aa')
benchmark(reverseVowels2, 'leetcode')
