/**
 * 
 * @param {*} f 
 * @param  {...any} params 
 * @returns 
 */
function benchmark(f, ...params) {
    const initialCpuUsage = totalCpuUseInSeconds()
    f(...params)
    const finalCpuUsage = totalCpuUseInSeconds()
    const stats = {
        function: f.name,
        params: Array.from(...params).join(''),
        cpuUsageInSeconds: finalCpuUsage - initialCpuUsage
    }
    console.table(stats)
    return stats
}

function totalCpuUseInSeconds() {
    return (process.cpuUsage().user + process.cpuUsage().system)/1000000
}

module.exports = { benchmark }
