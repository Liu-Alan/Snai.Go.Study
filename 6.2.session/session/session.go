package session

import (
	"fmt"
	"sync"
)

//session存储方式接口
type Provider interface {
	//初始化一个session,sid根据需要生成后传入
	SessionInit(sid string) (Session, error)
	//根据sid,获取session
	SessionRead(sid string) (Session, error)
	//销毁session
	SessionDestroy(sid string) error
	//回收
	SessionGC(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

type Manager struct {
	cookieName  string
	lock        sync.Mutex //互斥锁
	provider    Provider   //存储session方式
	maxLifeTime int64      //有效期
}

//实例化一个session管理器
func NewSessionManager(provideName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provide, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session:unknown provide %q", provideName)
	}
	return &Manager{cookieName: cookieName, provider: provide, maxLifeTime: maxLifeTime}, error
}

//注册 由实现Provider接口的结构体调用
func Register(name string, provide Provider) {
	if provide == nil {
		panic("session:Register provide is nil")
	}
	if _, ok := provides[name]; ok {
		panic("session:Register called twice for provide " + name)
	}
	provides[name] = provide
}

var provides = make(map[string]Provider)
