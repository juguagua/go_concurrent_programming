package sync_atomic

import (
	"sync/atomic"
	"unsafe"
)

// LKQueue lock free queue
type LKQueue struct {
	head unsafe.Pointer // head 是一个 dummy node
	tail unsafe.Pointer // tail 指向实际的队尾
}

// 队列通过链表实现
type node struct {
	value any
	next  unsafe.Pointer
}

// NewLKQueue new a lock free queue
func NewLKQueue() *LKQueue {
	n := unsafe.Pointer(&node{})
	return &LKQueue{head: n, tail: n}
}

func (q *LKQueue) Enqueue(v any) {
	n := &node{value: v}
	for {
		tail := load(&q.tail)
		next := load(&tail.next)
		if tail == load(&q.tail) { // 如果 tail 值没有变化
			if next == nil { // 还没有新的数据入队
				if cas(&tail.next, next, n) { // 增加到队尾
					cas(&q.tail, tail, n) // 入队成功，移动尾巴指针
					return
				}
			} else { // 已有新数据加到队列后面，移动尾指针
				cas(&q.tail, tail, next)
			}
		}
	}
}

func (q *LKQueue) Dequeue() any {
	for {
		head := load(&q.head)
		tail := load(&q.tail)
		next := load(&head.next)
		if head == load(&q.head) { // 如果 head 值没有变化
			if head == tail { // 如果 head 和 tail 一样
				if next == nil { // 同时 next 为 null 说明是空队列
					return nil
				}
				// next 不为 null 说明 只是 tail 指针没有调整位置，调整它
				cas(&q.tail, tail, next)
			} else {
				// 读取出队的数据
				v := next.value // 注意 head 是一个 dummy node，没有实际的值，实际的值在 next 上
				// 移动头指针
				if cas(&q.head, head, next) {
					return v // 返回数据，完成出队
				}

			}
		}
	}
}

// 将 unsafe.Pointer 原子地加载成 node
func load(p *unsafe.Pointer) (n *node) {
	return (*node)(atomic.LoadPointer(p))
}

// 封装 CAS，避免直接将 *node 转换成 unsafe.Pointer
func cas(p *unsafe.Pointer, old, new *node) (ok bool) {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}
