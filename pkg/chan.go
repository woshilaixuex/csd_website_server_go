package pkg

import (
	"sync"
)

/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description:
 * @Date: 2024-09-28 15:05
 */
const ClosedSignal = "channel closed"

type MssChan struct {
	Max        int
	isopen     bool
	sync.Mutex // 防止关闭时造成阻塞
	C          chan interface{}
}

func NewMssChan(max int) *MssChan {
	return &MssChan{
		Max:    max,
		isopen: false,
		C:      make(chan interface{}, max),
	}
}
func (m *MssChan) CheckIsOpen() bool {
	m.Lock()
	defer m.Unlock()
	return m.isopen
}
func (m *MssChan) Open() {
	m.Lock()
	defer m.Unlock()
	if !m.isopen {
		m.isopen = true
	}
}

// Close 关闭通道并标记为不可用
func (m *MssChan) Close() {
	m.Lock()
	defer m.Unlock()
	if m.isopen {
		m.C <- ClosedSignal // 发送关闭信息
		close(m.C)
		m.isopen = false
	}
}

// Send 向通道发送数据，如果通道已关闭则返回错误
func (m *MssChan) Send(data interface{}) {
	m.Lock()
	defer m.Unlock()
	if !m.isopen {
		return
	}
	m.C <- data
}
