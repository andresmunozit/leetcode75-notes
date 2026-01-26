const { strStr } = require('.');

describe('strStr', () => {
    const tests = [
        {
            name: 'needle found',
            neelde: 'sad',
            haystack: 'sadbutsad',
            expected: 0,
        },
        {
            name: 'needle not found',
            neelde: 'leeto',
            haystack: 'leetcode',
            expected: -1,
        },
        {
            name: 'needle equals to haystack',
            neelde: 'a',
            haystack: 'a',
            expected: 0,
        },
    ];

    tests.forEach((t) => {
        test(t.name, () => {
            expect(strStr(t.haystack, t.neelde)).toBe(t.expected);
        });
    });
});
