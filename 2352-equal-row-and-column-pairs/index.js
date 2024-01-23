/**
 * Counts the number of equal row-column pairs in a square matrix.
 * @param {number[][]} grid - The input square matrix.
 * @return {number} - The number of equal row-column pairs.
 */
var equalPairs = function(grid) {
    // Check that it's an square matrix
    if (grid.length !== grid[0].length) return 0

    // Initialize variables
    const rows = {}
    const columns = {}
    let equalPairsCount = 0

    // Populate the `rows` object, convert row contents to string for further comparison
    for (let i = 0; i < grid.length; i++) {
        rows[i] = grid[i].join()
    }

    // Populate the `columns` object, convert columns to string
    for (let i = 0; i < grid.length; i++) {
        for (let j = 0; j < grid.length; j++) {
            // We'll use commas to help us distinguish between columns and rows with similar
            // characters. For example if row = [11,1] and column = [1,11] are compare as strings,
            // without commas, the comparison will be '111' === '111' which will cause a false
            // positive
            columns[j] ? 
                columns[j] = columns[j] + ',' + grid[i][j].toString() :
                columns[j] = '' + grid[i][j].toString()
        }
    }

    // Compare rows with columns and increment the pair counter
    for (const row in rows) {
        for(const column in columns) {
            if (rows[row] === columns[column]) {
                equalPairsCount++
            }
        }
    }

    return equalPairsCount
};

const grid = [[3,2,1],[1,7,6],[2,7,7]]
var output = equalPairs(grid)

module.exports = equalPairs
