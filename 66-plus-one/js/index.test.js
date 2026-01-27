const { plusOne } = require('.');

describe('plusOne', () => {
    const tests = [
        {
            name: 'add one, no carry 1',
            digits: [1, 2, 3],
            expected: [1, 2, 4],
        },
        {
            name: 'add one, no carry 2',
            digits: [1, 2, 3],
            expected: [1, 2, 4],
        },
        {
            name: 'add one with carry, digits length equals to 1',
            digits: [9],
            expected: [1, 0],
        },
        {
            name: 'add one with carry, new digit at the end, digits length greater than 1',
            digits: [9, 9, 9],
            expected: [1, 0, 0, 0],
        },
        {
            name: 'add one with carry, digits length greater than 1',
            digits: [1, 5, 9],
            expected: [1, 6, 0],
        },
    ];

    tests.forEach((t) => {
        test(t.name, () => {
            expect(plusOne(t.digits)).toStrictEqual(t.expected);
        });
    });
});
