import Exceljs from 'exceljs'

export const autoWidthAction = (val) => { 
    if(val == null || val === ''){
        return 10
    }
    const str = val.toString()
    let width = 0
    for(let i = 0; i < str.length; i++){
        const charCode = str.charCodeAt(i)
        // 中文字符、中文标点等占2个单位宽度
        width += (charCode > 255) ? 2 : 1
    }
    return Math.max(width, 10)  // 最小宽度10
}
export const exportExcel = async ({column, data, fileName, autowidth, format='xlsx'}) => {
    
    // 🔴 关键：将 Proxy 数据转换为普通对象
    const plainData = data.map(row => {
        const plainRow = {};
        // 遍历列配置，确保每个字段都被提取
        column.forEach(col => {
            // 兼容 Proxy 和普通对象，用可选链防止报错
            plainRow[col.name] = row?.[col.name] ?? ''; 
        });
        return plainRow;
    });
    
    // 创建工作簿
    const workbook = new Exceljs.Workbook();
    workbook.creator = 'admin';
    workbook.title = fileName;
    workbook.created = new Date();
    workbook.modified = new Date();
    
    // 添加sheet页
    const worksheet = workbook.addWorksheet(fileName);
  
    // 配置列（含宽度自动适配）
    const columnName = [];
    column.forEach((item, index) => {
        const obj = {
            header: item.label,
            key: item.name,
            width: null
        };
        if (autowidth) {
            const maxArr = [autoWidthAction(item.label)]; // 假设 autoWidthAction 是计算宽度的工具函数
            plainData.forEach(ite => { // 注意：这里要用转换后的 plainData！
                const str = ite[item.name] || '';
                if (str) maxArr.push(autoWidthAction(str));
            });
            obj.width = Math.max(...maxArr) + 5;
        }
        columnName.push(obj);
    });
    worksheet.columns = columnName;
    
    // 🔴 关键：用转换后的 plainData 写入行
    worksheet.addRows(plainData); 
    
    // 生成文件
    const buffer = format === 'xlsx' 
        ? await workbook.xlsx.writeBuffer() 
        : await workbook.csv.writeBuffer();
    const blob = new Blob([buffer], { type: 'application/octet-stream' });
    
    // 下载文件
    if (window.navigator.msSaveOrOpenBlob) {
        navigator.msSaveOrOpenBlob(blob, `${fileName}.${format}`);
    } else {
        const link = document.createElement('a');
        link.href = window.URL.createObjectURL(blob);
        link.download = `${fileName}.${format}`;
        link.click();
        window.URL.revokeObjectURL(link.href);
    }
}