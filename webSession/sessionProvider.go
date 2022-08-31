package webSession

//session存储器
type SessionProvider interface {
	create(sessionId string, data map[string]interface{}) error //创建session存储空间
	get(sessionId, key string) (string, error)                  //读取session键值
	getAll(sessionId string) (map[string]string, error)         //读取session所有键值对
	set(sessionId, key string, value interface{}) error         //设置session键值
	destroy(sessionId string) error                             //销毁session
	gc(expire int64) error                                      //垃圾回收：删除过期session
}
