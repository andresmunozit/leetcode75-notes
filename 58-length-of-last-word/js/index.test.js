const { lengthOfLastWord } = require('.');

describe('lengthOfLastWord', () => {
    const tests = [
        {
            name: 'success case',
            s: 'Hello World',
            expected: 5,
        },
        {
            name: 'string padded with spaces',
            s: '   fly me   to   the moon  ',
            expected: 4,
        },
        {
            name: 'success case 2',
            s: 'luffy is still joyboy',
            expected: 6,
        },
    ];

    tests.forEach((t) => {
        test(t.name, () => {
            expect(lengthOfLastWord(t.s)).toBe(t.expected);
        });
    });
});
