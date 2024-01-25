const assert = require('assert')
const asteroidCollision = require('./index')

const tests = [
    {
        asteroids: [5,10,-5],
        expectedOutput: [5,10]
    },
    {
        asteroids: [8,-8],
        expectedOutput: []
    },
    {
        asteroids: [10,2,-5],
        expectedOutput: [10]
    },
]

for (const t of tests) {
    assert.strict.deepEqual(asteroidCollision(t.asteroids), t.expectedOutput)
}
