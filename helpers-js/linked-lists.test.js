const { linkedListToArray, arrayToLinkedList } = require('./linked-lists');

test('transforms linked list to array', () => {
    let head = {
        val: 1,
        next: null,
    };
    let tail = head;

    tail.next = {
        val: 3,
        next: null,
    };
    tail = tail.next;

    tail.next = {
        val: 2,
        next: null,
    };
    tail = tail.next;

    tail.next = {
        val: 8,
        next: null,
    };
    tail = tail.next;

    expect(linkedListToArray(head)).toStrictEqual([1, 3, 2, 8]);
});

describe('transforms array to linked list', () => {
    test('valid case', () => {
        const arr = [1, 8, 10, 0, 3];
        
        let head = arrayToLinkedList(arr);
        expect(head.val).toBe(1);
        expect(head.next.val).toBe(8);
    
        head = head.next;
        expect(head.next.val).toBe(10);
    
        head = head.next;
        expect(head.next.val).toBe(0);
    
        head = head.next;
        expect(head.next.val).toBe(3);
    
        head = head.next;
        expect(head.next).toBe(null);
    });

    test('null case', () => {
        expect(arrayToLinkedList(null)).toBe(null);
    });

    test('empty case', () => {
        expect(arrayToLinkedList([])).toBe(null);
    });
});
