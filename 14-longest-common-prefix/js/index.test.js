const { longestCommonPrefix } = require('.');

describe('longestCommonPrefix', () => {
    const tests = [
        {
            name: 'there is common prefix',
            strs: ['flower', 'flow', 'flight'],
            expected: 'fl',
        },
        {
            name: 'there is no common prefix',
            strs: ['dog', 'racecar', 'car'],
            expected: '',
        },
        {
            name: 'all the strings contain the prefix',
            strs: ['cat', 'cat', 'cat'],
            expected: 'cat',
        },
    ];

    tests.forEach((t) => {
        test(t.name, () => {
            expect(longestCommonPrefix(t.strs)).toEqual(t.expected);
        });
    });
});
