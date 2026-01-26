const { arrayToLinkedList } = require("../../helpers-js/linked-lists");
const { mergeTwoLists } = require(".");

const toList = arrayToLinkedList;

describe('mergeTwoLists', () => {
    const tests = [
        {
            name: 'success case',
            list1: toList([1, 2, 4]),
            list2: toList([1, 3, 4]),
            want: toList([1, 1, 2, 3, 4, 4]),
        },
        {
            name: 'both lists are empty',
            list1: toList([]),
            list2: toList([]),
            want: toList([]),
        },
        {
            name: 'list1 is empty',
            list1: toList([]),
            list2: toList([0]),
            want: toList([0]),
        },
        {
            name: 'list2 is empty',
            list1: toList([0]),
            list2: toList([]),
            want: toList([0]),
        },
        {
            name: 'list1 is shorter',
            list1: toList([20, 25]),
            list2: toList([15, 41, 45, 52]),
            want: toList([15, 20, 25, 41, 45, 52]),
        },        
        {
            name: 'list2 is shorter',
            list1: toList([12, 53, 54, 88]),
            list2: toList([0, 12]),
            want: toList([0, 12, 12, 53, 54, 88]),
        },
    ];

    tests.forEach((t) => {
        test(t.name, () => {
            expect(mergeTwoLists(t.list1, t.list2)).toStrictEqual(t.want);
        });
    })
})