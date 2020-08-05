create_iframe=(src="https://www.baidu.com")=>{
    z=document.createElement('iframe')
    z.src=src
    z.width='500px'
    z.height='500px'
    z.onload=console.log
    document.body.appendChild(z)
    return z
}

window.frames

z.contentWindow.self
z.contentWindow.self.innerHTML

z.contentWindow.postMessage('zzz',location.href)
z.contentWindow.postMessage('zzz',"*")
z.contentWindow.postMessage('zzz',z.src)

window.parent.postMessage('dddd',z.src)




