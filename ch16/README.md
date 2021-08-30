## 模板  

text/template包接受一个模板字符串（或从文件加载的模板），并将数据插入其中。  

html/template包的工作方式与text/template类似，只是它还提供了使用HTML所需的安全保护。  

## 要点  

■　模板字符串包含将逐字输出的文本。在此文本中，可以插入含有将被计算的简单代码的各种action。可以使用action将数据插入模板文本中。  
■　Template值的Execute方法接受一个满足io.Writer接口的值，以及可以在模板的action中访问的数据值。  
■　模板action可以引用传递给使用{{.}}的Execute的数据值，简称“点”。点的值可以在模板的不同上下文中变化。  
■　在{{if}}action与其对应的{{end}}标记之间的模板部分，只有在某个条件为真时才会被包含。  
■　在{{range}}action与其对应的{{end}}标记之间的模板部分，将对数组、切片、映射或channel中的每个值重复。该部分中的任何操作也将被重复。  
■　在{{range}}节中，点的值将被更新为指向正在处理的集合中的当前元素。  
■　如果点指向一个struct值，则可以用{{.FieldName}}插入该struct中的字段值。  
■　当浏览器需要从服务器获取数据时，通常使用HTTP GET请求。  
■　当浏览器需要向服务器提交新数据时，将使用HTTP POST请求。  
■　可以使用http.Request值的FormValue方法访问来自请求的表单数据。  
■　可以使用http.Redirect函数来引导浏览器请求不同的路径。  