package znet

import (
	"../ziface"
	"fmt"
	"net"
	"time"
)

//iServer接口实现，定义一个server服务类
type Server struct {
	// 服务器的名称
	Name string
	// tcp4 or other
	IPVersion string
	// 服务器绑定的IP地址
	IP string
	//服务绑定的端口
	Port int
}

//实现ziface.IServer里的全部接口

//开启网络服务
func (s *Server) Start() {
	fmt.Printf("[START] Server listenner at IP:%s, Port:%d, is starting\n\n", s.IP, s.Port)
	//开始一个go去做服务端的Linster业务
	go func() {
		//1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error:", err)
			return
		}
		//2 监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		//监听成功
		fmt.Println("start Zinx server  ", s.Name, " succ, now listenning...")
		//3 启动server网络连接业务
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			//3.2 TODO Server.Start()设置服务器最大连接控制，如果超过最大连接，那么关闭新的连接
			//3.3 TODO Server。Start()处理该新连接请求业务的方法，此时应该有handler和conn是绑定的
			//这里使用一个最大512字节的回显服务
			go func(){
				// 不断的循环从和护短获取数据
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err", err)
						continue
					}
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err", err)
						continue
					}
				}
			}()
		}

	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)
}

func (s *Server) Serve() {
	s.Start()
	for {
		time.Sleep(10 * time.Second)
	}
}

func NewServer (name string) ziface.IServer {
	s := &Server{
		Name: name,
		IPVersion: "tcp4",
		IP: "0.0.0.0",
		Port: 7777,
	}
	return s
}