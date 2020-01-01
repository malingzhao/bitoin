## 创建客户端
1.  创建cli的结构体
 
 2. 测试
  1. 执行
     go build  -o bc.exe  main.go 
  2. 添加区块
     bc.exe addblock -data "a send 100 BTC to b"
     bc.exe addblock -data "b send 200 BTC to c"
 
  3. 将所有的区块打印出来
    bc.exe printchain 
    

  