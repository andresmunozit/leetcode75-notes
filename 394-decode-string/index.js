/**
 * 
 * @param {*} s 
 * @returns 
 */
const decodeString = s => {
    const stack = []
    stack.peek = () => stack[stack.length - 1]

    for (let c of s) {
        // Everything will be pushed to stack, until a ']' char is found
        if (c !== ']') {
            stack.push(c)
            continue
        }

        // ']' found, let's obtain the substring that will be multiplied and added to the stack
        let subStr = ''
        while (stack.peek() !== '[') {
            subStr = stack.pop() + subStr
        }

        // Eliminate '[' from the stack before to proceed
        stack.pop()

        // '[' found, let's obtain the multiplier
        let multiplierStr = ''
        while (!Number.isNaN(Number(stack.peek()))) {
            multiplierStr = stack.pop() + multiplierStr
        }

        const multiplier = Number(multiplierStr)

        // Now let's muliply add the str to the stack
        for (let i = 0; i < multiplier; i++) {
            stack.push(subStr)
        }
    }

    delete stack.peek

    return stack.join('')
};

var output = decodeString('3[a]2[bc]')

module.exports = decodeString
