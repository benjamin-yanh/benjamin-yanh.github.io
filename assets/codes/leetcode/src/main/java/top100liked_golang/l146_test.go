package top100liked_golang

import (
	"fmt"
	"testing"
)

type Entry struct {
	key, value int
	prev, next *Entry
}

type LRUCache struct {
	capacity  int
	dummy     *Entry // 哨兵节点
	keyToNode map[int]*Entry
}

func Constructor(capacity int) LRUCache {
	dummy := &Entry{}
	dummy.prev = dummy
	dummy.next = dummy
	return LRUCache{
		capacity:  capacity,
		dummy:     dummy,
		keyToNode: map[int]*Entry{},
	}
}

// 删除一个节点（抽出一本书）
func (c *LRUCache) remove(x *Entry) {
	x.prev.next = x.next
	x.next.prev = x.prev
}

// 在链表头添加一个节点（把一本书放到最上面）
func (c *LRUCache) pushFront(x *Entry) {
	x.prev = c.dummy
	x.next = c.dummy.next
	x.prev.next = x
	x.next.prev = x
}

// 获取 key 对应的节点，同时把该节点移到链表头部
func (c *LRUCache) getNode(key int) *Entry {
	node := c.keyToNode[key]
	if node == nil { // 没有这本书
		return nil
	}
	c.remove(node)    // 把这本书抽出来
	c.pushFront(node) // 放到最上面
	return node
}

func (c *LRUCache) Get(key int) int {
	node := c.getNode(key) // getNode 会把对应节点移到链表头部
	if node == nil {
		return -1
	}
	return node.value
}

func (c *LRUCache) Put(key, value int) {
	node := c.getNode(key) // getNode 会把对应节点移到链表头部
	if node != nil {       // 有这本书
		node.value = value // 更新 value
		return
	}
	node = &Entry{key: key, value: value} // 新书
	c.keyToNode[key] = node
	c.pushFront(node)                  // 放到最上面
	if len(c.keyToNode) > c.capacity { // 书太多了
		backNode := c.dummy.prev
		delete(c.keyToNode, backNode.key)
		c.remove(backNode) // 去掉最后一本书
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func Test_LRUCache(t *testing.T) {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))

}
