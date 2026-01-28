/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function(n) {
    // Known cases:
    // step 0 -> 1 way
    // step 1 -> 1 way
    // step 2 -> 2 ways
    const ways = [1, 1, 2];

    // steps = [0, 1, 2, 3, 4]
    // ways  = [1, 1, 2, ?, ?]

    // To reach step 3:
    // - come from step 2 (1 step)
    // - come from step 1 (2 steps)
    // ways[3] = ways[2] + ways[1] = 2 + 1 = 3

    // To reach step 4:
    // - come from step 3 (1 step)
    // - come from step 2 (2 steps)
    // paths to step 4 come from step 3 or step 2, so we add their counts
    // ways[4] = ways[3] + ways[2] = 3 + 2 = 5

    // We count ways, not steps. Same rule for every next step.
    for (let i = 3; i <= n; i++) {
        ways[i] = ways[i-1] + ways[i-2];
    }

    return ways[n];
};
