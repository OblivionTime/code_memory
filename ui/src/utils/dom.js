import { ElMessage } from 'element-plus'
import useClipboard from "vue-clipboard3";
export function addLine(block) {
    //添加行号
    // 1.创建ul节点
    let ul = document.createElement("ul");
    if (!block.textContent) {
        return
    }
    // 2.根据换行符获取行数，根据获取的行数生成行号
    let rowCount = block.outerHTML.split('\n').length;
    for (let i = 1; i <= rowCount; i++) {
        //创建li节点，创建文本节点
        let li = document.createElement("li")
        let text = document.createTextNode(i)
        //为li追加文本节点，将li加入ul
        li.appendChild(text)
        ul.appendChild(li)
    }
    // 3.给行号添加类名
    ul.className = 'pre-numbering'
    // 4.将ul节点加到 代码块。这个注意根据实际情况决定放在什么位置。是父节点，本节点还是子节点上
    block.appendChild(ul)
}
const { toClipboard } = useClipboard();
//复制代码
export async function copyCode(code) {
    try {
        await toClipboard(code);  //实现复制
        ElMessage({
            message: '复制成功',
            type: 'success',
        })
    } catch (e) {
        ElMessage({
            message: '复制失败',
            type: 'warning',
        })
    }
}