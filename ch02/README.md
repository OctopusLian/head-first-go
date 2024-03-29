## 要点  

```
■  方法是一种与给定类型的值相关联的函数。
■　Go将从//标记到行尾的所有内容都视为注释，并忽略它。
■　多行注释以/*开头，以*/结尾。中间的所有内容，包括换行符，都将被忽略。
■　传统的做法是在每个程序的顶部都加上一条注释，解释它是做什么的。
■　与大多数编程语言不同，Go允许从函数调用或方法调用返回多个值。
■　多个返回值的一个常见用法是返回函数的主要结果，然后返回第二个值，指示是否有错误。
■　若要在不使用值的情况下丢弃它，请使用_空白标识符。空白标识符可以用于任何变量的任何赋值语句中。
■　避免给变量取与类型、函数或包相同的名称，这会使变量遮盖（覆盖）具有相同名称的内容。
■　函数、条件和循环都有出现在{}大括号中的代码块。
■　文件和包的代码不出现在{}大括号中，但是它们也包含块。
■　变量的作用域仅限于它被定义的块内以及嵌套在该块中的所有块。
■　除了名称之外，导入包时可能还需要包的导入路径。
■　continue关键字跳转到循环的下一个迭代。
■　break关键字完全退出循环。
```