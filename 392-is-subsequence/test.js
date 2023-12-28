const assert = require('assert')
const isSubsequence = require('./index')

const tt = [
    {
        s: 'abc',
        t: 'ahbgdc',
        expected: true,
    },
    {
        s: 'axc',
        t: 'ahbgdc',
        expected: false,
    }
]

for (const t of tt) {
    assert.strict.equal(isSubsequence(t.s, t.t), t.expected)
}
