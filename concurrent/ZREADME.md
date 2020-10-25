#go语言 高并发练习总结
> 主协：特点
+ 1.主协成死掉，子协成陪着一起死 test01
+ 2.runtime.Gosched() 让出协成为日成 降低优先级 test02
+ 3.查看cpu的核心数和设置cpu的核心数 test03
    ```
        var gomaxprocs int
         runtime.NumCPU() //查看cpu的核心数
        gomaxprocs = runtime.GOMAXPROCS(2)// 返回先前的cpu的核心数为当前的电脑配置
        fmt.Println(gomaxprocs)
        gomaxprocs = runtime.GOMAXPROCS(1)//2 返回先前的cpu的核心数为2
        fmt.Println(gomaxprocs)
        gomaxprocs = runtime.GOMAXPROCS(3)//1 返回先前的cpu的核心数为1
        fmt.Println(gomaxprocs)
    ```
 + 4.协成自杀,其他协成不受影响 runtime.Goexit() test04
 + 5.管道机制做通讯  
     ```
        管道必须读才能写的进去，没人读，一直处于阻塞状态, 
        管道为零的管道可以读，不可以写  
        chan<- 只写  <-chan只读
        close()  关闭管道：注意下一：
       1.关闭为未初始化、重复关闭一个管道、向关闭的管道发送消息、的管道会产生异常
       2.读关闭的管道 永不阻塞，切有广播的效果 都可以false判断
        mychan := make(chan int, 8)
            mychan<-1000
            fmt.Println(len(mychan)) //1   长度
            fmt.Println(cap(mychan))  //8  容量
        
    ```
+ 6 生产消费者test06 -09
   ```
        //生产者每秒生成一件商品，并通知物流公司取货
        //物流公司将商品运输到超市
        //消费者阻塞等待从商品消费
        //消费10轮就主协成结束
    ```
+ 7.控制并发数在10范围内
+ 8.select 的用法 
      ```
       flag://跳出标记位置
        for  {
            select {
            case chaA<-12:
                fmt.Println("随机执行任务,选择了我A")
                time.Sleep(time.Second)
            case <-chaB:
                fmt.Println("随机执行任务,选择了我B")
                time.Sleep(time.Second)
            case <-chaC:
                fmt.Println("随机执行任务,选择了我C")
                time.Sleep(time.Second)
            default:
                fmt.Println("没有任务,情况下我来执行    *^-^*   ")
                time.Sleep(time.Second)
                break flag//跳出标记位置
            }
        }
        ```
+ 9.timer定时器   定时的和周期的两种
    ```
    //定时器 一次性的
        timer := time.NewTimer(2*time.Second)   等价于  time.After(time.Second)
        fmt.Println(time.Now())
        //停止防止计时器启动。
        //如果调用停止计时器，则返回true；如果计时器已经停止，则返回false
        //过期或被停止
        //timer.Stop()//停止
        timer.Reset(time.Second) //重置
        x:=<-timer.C
        fmt.Println(x)
        //秒 周期性
        ticker := time.NewTicker(2 * time.Second)
    
        for  {
            fmt.Println(<-ticker.C)
        }
    ```
+ 10.等待组
     ```
        var wg sync.WaitGroup
            wg.Add(1) //添加一个协成
            wg.Done() //注销一个协成
            wg.Wait() //阻塞等待协成数为零停止等待
     ```
+ 11.读写锁(数据库 一写多读  test15)
  ```
  var rw sync.RWMutex
  	rw.RLock() //读锁
  	rw.RUnlock() //释放读锁
  	rw.Lock() 	//写锁
  	rw.Unlock() //释放读锁
  
  ```
  
+ 12. sync.Once  once.Do 只执行一次  test16

+ 13 cond := sync.NewCond(&sync.Mutex{})  条件变量控制(变量通讯)  test18 发送和接收监听值的变化
+ 14 "sync/atomic“   	atomic.AddInt64()  //原子操作变量
 



 




