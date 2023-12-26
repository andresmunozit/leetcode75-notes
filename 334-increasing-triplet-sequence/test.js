const assert = require('assert')
const increasingTriplet = require('./index')

assert.strict.equal(increasingTriplet([1,2,3,4,5]), true)
assert.strict.equal(increasingTriplet([5,4,3,2,1]), false)
assert.strict.equal(increasingTriplet([2,1,5,0,4,6]), true)
assert.strict.equal(increasingTriplet([20,100,10,12,5,13]), true)
