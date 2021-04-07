package LRU

import "fmt"

type LruNode struct {
	Key int
	Value interface{}
	Pre   *LruNode
	Next  *LruNode
}

type LruCache struct {
	Head *LruNode
	Tail *LruNode
	L map[int]*LruNode
	Size  int
	MaxSize int
}

func NewLruCache(maxSize int)*LruCache{
	return &LruCache{
		Head:nil,
		Tail:nil,
		L:       make(map[int] *LruNode),
		Size:    0,
		MaxSize: maxSize,
	}
}

// 查询操作
func (p *LruCache)Get(key int) interface{} {
	//如果有这个元素，把这个元素移动到链表头部，如果不存在返回nil
	if v,ok := p.L[key];ok{
		if v != p.Head {
			oldNext := v.Next
			oldPre := v.Pre

			oldPre.Next = oldNext
			if oldNext != nil {
				oldNext.Pre = oldPre
			}
			p.Head.Pre = v
			v.Next = p.Head
			v.Pre = nil
			p.Head = v

		}
		return v.Value
	}else{
		return nil
	}
}

// 插入操作,也可能是更新
func (p *LruCache)Put(key int, value interface{}) {
	// 先查看key 是否存在,若存在则更新
	if v,ok := p.L[key];ok{
		v.Value = value
		if v != p.Head {
			oldNext := v.Next
			oldPre := v.Pre

			oldPre.Next = oldNext
			if oldNext != nil {
				oldNext.Pre = oldPre
			}
			p.Head.Pre = v
			v.Next = p.Head
			v.Pre = nil
			p.Head = v
		}
	}else{
		newNode := &LruNode{
			Key:key,
			Value:value,
			Pre:nil,
			Next:nil,
		}
		//若不存在 ，先查看是否已经满了
		if p.Size == p.MaxSize { //触发lru 淘汰机制,删除尾节点
			oldPre := p.Tail.Pre
			oldPre.Next = nil
			delete(p.L,p.Tail.Key)
			p.Tail = oldPre

			p.Head.Pre = newNode
			newNode.Next = p.Head
			p.Head = newNode

		}else{
			if p.Size == 0 {
				p.Head = newNode
				p.Tail = newNode
			}else{
				p.Head.Pre = newNode
				newNode.Next = p.Head
				p.Head = newNode
			}
			p.Size++
		}
		p.L[key] = newNode
	}
}


func (p *LruCache)Print() {
	e := p.Head
	for e != nil {
		fmt.Printf("%d ",e.Key)
		e = e.Next
	}
	fmt.Println("")
}
