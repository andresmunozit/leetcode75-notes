const { searchInsert } = require('.');

describe('searchInsert', () => {
    const tests = [
        {
            name: 'nums length is zero',
            nums: [],
            target: 5,
            expected: 0,
        },
        {
            name: 'target is less than min',
            nums: [1, 3, 5, 6],
            target: 0,
            expected: 0,
        },
        {
            name: 'target is greater than max',
            nums: [1, 3, 5, 6],
            target: 7,
            expected: 4,
        },
        {
            name: 'target is greater than max by two',
            nums: [2, 3, 4, 7, 8, 9],
            target: 11,
            expected: 6,
        },
        {
            name: 'target is present',
            nums: [1, 3, 5, 6],
            target: 5,
            expected: 2,
        },
        {
            name: 'target is between two nums',
            nums: [1, 3, 5, 6],
            target: 2,
            expected: 1,
        },
        {
            name: 'nums length is two',
            nums: [1, 3],
            target: 2,
            expected: 1,
        },
    ];

    tests.forEach((t) => {
        test(t.name, () => {
            expect(searchInsert(t.nums, t.target)).toBe(t.expected);
        });
    });
});
