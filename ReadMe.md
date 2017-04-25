
最近有频繁从mac到远程vps传输大量小文件到vps的需要，然并卵，本地主机到vps的网络延迟很大(ping 160ms)，同时存在连接不稳定的问题，居然没有找到一个合适的同步工具来完成这件事，于是用Golang 写了一个文件同步工具.

尝试过SCP，主要存在两个问题：
1. 文件同步过慢
2. 拷贝时间非常长，中间连接断开又要重新拷贝，基本无法完成一次完整的拷贝；

###说明：

1. 支持本地与远程目录的同步；双方使用tcp通信，本地作为client端，远程作为server端；
2. 支持大文件过滤，超过10M的文件不做同步；
3. 支持文件比较；忽略本地和远程目录中相同的文件；
4. 传输速率高；数据帧使用IO复用全异步转发，控制通道和数据传输通道分开，数据通道可以同时建N条，提升传输速率；
5. 可靠性保证： 对于同一个文件使用同一条tcp连接，解决接收端文件包可能错乱的问题；
6. 通信协议采用PB协议；
7. 连接断开，支持自动重连；

###同步流程
1. client 遍历source Dir中的文件，并计算其hash值，通道控制通道发送给Server端；
2. Server根据文件名寻找dest Dir中的文件，计算hash与收到的file/hash比较，如果不同，通过控制通道发送文件请求给client
3. client收到文件请求后，对文件内容进行分包，通过数据通道发送给server
4. Server端收到数据分片，组合保存到文件；

###TODO
2017-04-24

1. 加上连接心跳
2. 支持配置文件中获取参数； (done)
3. 支持从args中获取同步参数；
4. 优化打印输出； （done）

