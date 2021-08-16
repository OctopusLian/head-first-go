## 要点  

■　所有Go程序都至少有一个goroutine：程序启动时调用main函数的那个goroutine。  
■　Go程序在main goroutine停止时结束，即使其他goroutine尚未完成其工作。  
■　time.Sleep函数将当前goroutine暂停一段时间。  
■　Go不保证何时在goroutine之间切换，或者它将持续运行一个goroutine多长时间。这允许goroutine更高效地运行，但这意味着你不能指望按特定顺序执行操作。  
■　函数返回值不能在go语句中使用，部分原因是当调用函数试图使用它时，返回值还没有准备好。  
■　如果需要goroutine中的值，则需要将其传递给一个channel，以便将该值发送回来。  
■　channel是通过调用内置的make函数创建的。  
■　每个channel只携带一种特定类型的值，在创建channel时指定该类型。  
■　使用<-运算符将值发送给channel：  
■　<-运算符也用于从channel接收值：  