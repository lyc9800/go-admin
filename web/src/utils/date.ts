/* 
    @param time 需要转换的时间
    @param fmt 需要转换的格式 如 yyyy-MM-dd yyyy-MM-dd HH:mm:ss
    @return 格式化后的时间字符串
*/
export function formatTime(time: any, fmt: string): string {
    if (!time) return ''
    
    const date = new Date(time)
    
    const o: { [key: string]: number } = {
        'M+': date.getMonth() + 1,                     // 月份
        'd+': date.getDate(),                          // 日
        'H+': date.getHours(),                         // 小时
        'h+': date.getHours() % 12 || 12,              // 12小时制小时
        'm+': date.getMinutes(),                       // 分
        's+': date.getSeconds(),                       // 秒
        'q+': Math.floor((date.getMonth() + 3) / 3),   // 季度
        'S': date.getMilliseconds()                    // 毫秒
    }
    
    // 处理年份
    if (/(y+)/.test(fmt)) {
        const yearMatch = fmt.match(/(y+)/)
        if (yearMatch) {
            const yearStr = date.getFullYear().toString()
            fmt = fmt.replace(yearMatch[0], yearStr.substr(4 - yearMatch[0].length))
        }
    }
    
    // 遍历替换其他时间单位
    for (const k in o) {
        const pattern = new RegExp(`(${k})`)
        if (pattern.test(fmt)) {
            const match = fmt.match(pattern)
            if (match) {
                let str = o[k].toString()
                // 如果是两位数格式（如 MM、dd、HH 等）
                if (match[0].length > 1) {
                    str = str.padStart(2, '0')
                }
                // 如果是三位数格式（如 SSS）
                if (match[0].length > 2) {
                    str = str.padStart(3, '0')
                }
                fmt = fmt.replace(match[0], str)
            }
        }
    }
    
    return fmt
}