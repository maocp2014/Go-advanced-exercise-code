for-select也是使用频率很高的结构，select提供了多路复用的能力，所以for-select可以让函数具有持续多路处理多个channel的能力。但select没有感知channel的关闭，这引出了2个问题：

继续在关闭的通道上读，会读到通道传输数据类型的零值，如果是指针类型，读到nil，继续处理还会产生nil。

继续在关闭的通道上写，将会panic。

问题2可以这样解决，通道只由发送方关闭，接收方不可关闭，即某个写通道只由使用该select的协程关闭，select中就不存在继续在关闭的通道上写数据的问题。

问题1可以使用,ok来检测通道的关闭，使用情况有2种。

第一种：如果某个通道关闭后，需要退出协程，直接return即可。示例代码中，该协程需要从in通道读数据，还需要定时打印已经处理的数量，有2件事要做，所有不能使用for-range，需要使用for-select，当in关闭时，ok=false，我们直接返回。

 1 go func() {
 2    // in for-select using ok to exit goroutine
 3    for {
 4        select {
 5        case x, ok := <-in:
 6            if !ok {
 7                return
 8            }
 9            fmt.Printf("Process %d\n", x)
10            processedCnt++
11        case <-t.C:
12            fmt.Printf("Working, processedCnt = %d\n", processedCnt)
13        }
14    }
15}()

第二种：如果某个通道关闭了，不再处理该通道，而是继续处理其他case，退出是等待所有的可读通道关闭。我们需要使用select的一个特征：select不会在nil的通道上进行等待。这种情况，把只读通道设置为nil即可解决。
 1 go func() {
 2    // in for-select using ok to exit goroutine
 3    for {
 4        select {
 5        case x, ok := <-in1:
 6            if !ok {
 7                in1 = nil
 8            }
 9            // Process
10        case y, ok := <-in2:
11            if !ok {
12                in2 = nil
13            }
14            // Process
15        case <-t.C:
16            fmt.Printf("Working, processedCnt = %d\n", processedCnt)
17        }
18
19        // If both in channel are closed, goroutine exit
20        if in1 == nil && in2 == nil {
21            return
22        }
23    }
24}()