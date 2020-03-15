package memory

import (
	"container/list"
	"sync"
	"time"

	"Snai.Go.Study/6.2.session/session"
)

var pder = &FromMemory{list: list.New()}

func init() {
	pder.sessions = make(map[string]*list.Element, 0)
	//注册memory调用的时候一定要一致
	session.Register("memory", pder)
}

//session实现
type SessionStore struct {
	sid              string                      //session id 唯一标识
	LastAccessedTime time.Time                   //最后访问时间
	value            map[interface{}]interface{} //session 里面存储的值
}

//设置
func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}

//获取session
func (st *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
	return nil
}

//删除
func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) SessionID() string {
	return st.sid
}

//session内存实现
type FromMemory struct {
	lock     sync.Mutex               //用来锁
	sessions map[string]*list.Element //用来存储在内存
	list     *list.List               //用来做gc
}

func (fromMemory *FromMemory) SessionInit(sid string) (session.Session, error) {
	fromMemory.lock.Lock()
	defer fromMemory.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, LastAccessedTime: time.Now(), value: v}
	element := fromMemory.list.PushBack(newsess)
	fromMemory.sessions[sid] = element
	return newsess, nil
}

func (fromMenory *FromMemory) SessionRead(sid string) (session.Session, error) {
	if element, ok := fromMenory.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := fromMenory.SessionInit(sid)
		return sess, err
	}
	return nil, nil
}

func (fromMemory *FromMemory) SessionDestroy(sid string) error {
	if element, ok := frommemory.sessions[sid]; ok {
		delete(frommemory.sessions, sid)
		frommemory.list.Remove(element)
		return nil
	}
	return nil
}

func (fromMemory *FromMemory) SessionGC(maxLifeTime int64) {
	fromMemory.lock.Lock()
	defer fromMemory.lock.Unlock()
	for {
		element := fromMemory.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).LastAccessedTime.Unix() + maxLifeTime) < time.Now().Unix() {
			fromMemory.list.Remove(element)
			delete(fromMemory.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (fromMemory *FromMemory) SessionUpdate(sid string) error {
	fromMemory.lock.Lock()
	defer fromMemory.lock.Unlock()
	if element, ok := fromMemory.sessions[sid]; ok {
		element.Value.(*SessionStore).LastAccessedTime = time.Now()
		fromMemory.list.MoveToFront(element)
		return nil
	}
	return nil
}
