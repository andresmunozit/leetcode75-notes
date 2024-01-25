const assert = require('assert')
const decodeString = require('./index')

const tests = [
    {
        s: '3[a]2[bc]',
        output: 'aaabcbc'
    },
    {
        s: '3[a2[c]]',
        output: 'accaccacc'
    },
    {
        s: '2[abc]3[cd]ef',
        output: 'abcabccdcdcdef'
    }
]

for (const t of tests) {
    assert(decodeString(t.s), t.output)
}
