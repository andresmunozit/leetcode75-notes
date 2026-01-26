function linkedListToArray(linkedList) {
    const arr = [];

    let head = linkedList;

    while (head !== null) {
        arr.push(head.val);
        head = head.next;
    }

    return arr;
}

function arrayToLinkedList(arr) {
    if (arr === null) {
        return null;
    }

    const dummy = {
        val: 0,
        next: null,
    };

    let tail = dummy;

    for (let i = 0; i < arr.length; i++) {
        tail.next = {
            val: arr[i],
            next: null,
        };

        tail = tail.next;
    }

    return dummy.next;
}

module.exports = {
    linkedListToArray,
    arrayToLinkedList,
}
