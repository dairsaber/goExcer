# 笔记
## runtime 下面的一些函数
-  `runtime.GoExit()` 退出正个go程；
- `runtime.GOMAXPROCS()` 调用cpu的核心个数；

## channel 通道
 ### `channel`  声明

    ```go
       currentChannel := make (chan type,capcity)
       // 其中type就是这个channel传递的数据类型
       // capcity 就是缓冲区的大小 capcity为0或者不传则表示没有缓冲区
    ```

- 特性：channel是具有读写两端的一个通道；当任何一端没有处理（未读/未写）都会造成阻塞 利用这点可以实现数据的顺序执行；其实就相当于加锁和解锁的过程