/**
 * @param {number[]} asteroids
 * @return {number[]}
 */
var asteroidCollision = function(asteroids) {
    const asteroidsStack = [asteroids[0]];
    asteroidsStack.peek = () => asteroidsStack[asteroidsStack.length - 1];

    let i = 1;
    while (i < asteroids.length) {
        // Detect collisions
        // If two asteroids meet, the smaller one (by absolute value) will explode
        if (asteroidsStack.peek() > 0 && asteroids[i] < 0) {
            // Iterate until all collisions are resolved
            while (asteroidsStack.peek() > 0 && asteroids[i] < 0) {
                // If both are the same size (by absolute value), both will explode
                if (asteroidsStack.peek() + asteroids[i] === 0) {
                    asteroidsStack.pop();
                    i++;
                // If the incoming asteroid is larger (moving to the left), remove the top of the stack
                } else if (asteroidsStack.peek() + asteroids[i] < 0) {
                    asteroidsStack.pop();
                // If the asteroid in the stack is larger (moving to the right), ignore the incoming asteroid
                } else {
                    i++;
                }
            }
        // If there's no collision, add the asteroid to the stack
        } else {
            asteroidsStack.push(asteroids[i]);
            i++;
        }
    }

    delete asteroidsStack.peek
    
    return asteroidsStack;
};

module.exports = asteroidCollision
