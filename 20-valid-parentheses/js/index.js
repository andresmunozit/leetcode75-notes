/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
    // use a stack
    // if the next char is a closing char, check the last element of the stack and if it's a corresponding
    // opening char, delete it from the stack
    let stack = [];

    for (let i = 0; i < s.length; i++) {
        if (s[i] === '(' || s[i] === '{' || s[i] === '[') {
            stack.push(s[i]);
            continue;
        }

        // If a closing char has not a corresponding opening character return false
        if (s[i] === ')' && stack[stack.length - 1] === '(') {
            stack.pop();
        } else if (s[i] === ']' && stack[stack.length - 1] === '[') {
            stack.pop();
        } else if (s[i] === '}' && stack[stack.length - 1] === '{') {
            stack.pop();
        } else {
            return false;
        }
    }

    return stack.length === 0;
};

console.log(isValid('()'));
console.log(isValid('()[]{}'));
console.log(isValid('(]'));
console.log(isValid('([])'));
console.log(isValid('([)]'));
// true
// true
// false
// true
// false