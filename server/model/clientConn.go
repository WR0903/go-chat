package model

import (
	"net"
)

type ClientConn struct{}

type ConnInfo struct {
	// 客户端连接信息数据结构
	Conn     net.Conn
	UserName string
}

var ClientConnsMap map[int]ConnInfo

func init() {
	ClientConnsMap = make(map[int]ConnInfo)
}

func (cc ClientConn) Save(userID int, name string, userConn net.Conn) {
	ClientConnsMap[userID] = ConnInfo{userConn, name} // 将客户端连接信息保存下来
}

func (cc ClientConn) Del(userConn net.Conn) {
	for id, connInfo := range ClientConnsMap { // 删除某客户端连接信息
		if userConn == connInfo.Conn {
			delete(ClientConnsMap, id)
		}
	}
}

func (cc ClientConn) SearchByUserName(userName string) (connInfo net.Conn, err error) {
	// 通过username获得客户端连接信息
	user, err := CurrentUserDao.GetUserByUserName(userName)
	if err != nil {
		return
	}

	connInfo = ClientConnsMap[user.ID].Conn
	return
}
