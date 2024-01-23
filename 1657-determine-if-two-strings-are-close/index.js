/**
 * @param {string} word1
 * @param {string} word2
 * @return {boolean}
 */
var closeStrings = function(word1, word2) {
    // Edge case: different lengths
    if (word1.length !== word2.length) {
        return false
    }

    // Obtain letter maps and sets
    const word1FreqMap = {}
    const word2FreqMap = {}
    const word1LetterSet = new Set()
    const word2LetterSet = new Set()
    let word1Freqs = []
    let word2Freqs = []

    for (let i = 0; i < word1.length; i++) {
        word1FreqMap.hasOwnProperty(word1[i]) ? word1FreqMap[word1[i]]++ : word1FreqMap[word1[i]] = 1
        word2FreqMap.hasOwnProperty(word2[i]) ? word2FreqMap[word2[i]]++ : word2FreqMap[word2[i]] = 1
    }

    // Let's create the sets
    for (const key in word1FreqMap) {
        word1LetterSet.add(key)
        word1Freqs.push(word1FreqMap[key])
    }
    for (const key in word2FreqMap) {
        word2LetterSet.add(key)
        word2Freqs.push(word2FreqMap[key])
    }

    // Let's compare sizes
    if (word1LetterSet.size !== word2LetterSet.size) {
        return false
    }

    // Let's iterate and compare letters and frequencies
    word1LetterSet.forEach((letter) => {
        if (!word2LetterSet.has(letter)) return false
        word2LetterSet.delete(letter)
    })

    if (word2LetterSet.size !== 0) return false

    word1Freqs = word1Freqs.sort()
    word2Freqs = word2Freqs.sort()

    // There should be the same frequencies in order of two strings to be close
    for (let i = 0; i < word1Freqs.length; i++) {
        if (word1Freqs[i] !== word2Freqs[i]) return false
    }

    return true
};

var output = closeStrings('abbzzca', 'babzzcz')

module.exports = closeStrings
