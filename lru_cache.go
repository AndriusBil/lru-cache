package lru_cache

type LinkedNode struct {
	value interface{}
	key   interface{}
	next  *LinkedNode
	prev  *LinkedNode
}

type LRUCache struct {
	store map[interface{}]*LinkedNode
	cap   int
	size  int
	head  *LinkedNode
	tail  *LinkedNode
}

// New creates a new LRUCache.
// The capacity variable determines how many items can be stored.
func New(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		size:  0,
		store: map[interface{}]*LinkedNode{},
	}
}

func (c *LRUCache) push(item *LinkedNode) {
	if c.size == 0 {
		c.head = item
		c.tail = item
		return
	}

	item.prev = c.head
	c.head.next = item
	c.head = item
}

func (c *LRUCache) promote(item *LinkedNode) {
	if item == c.head {
		return
	}

	if item == c.tail {
		c.tail = item.next
		c.tail.prev = nil
		c.push(item)

		return
	}

	if item.prev != nil && item.next != nil {
		item.prev.next = item.next
		item.next.prev = item.prev
	}

	c.push(item)
}

// Get returns stored value by provided variable key.
// Requested value is pushed to the head of the cache.
func (c *LRUCache) Get(key interface{}) interface{} {
	item, ok := c.store[key]

	if !ok {
		return nil
	}

	c.promote(item)

	return item.value
}

func (c *LRUCache) evict() {
	if c.tail == nil {
		return
	}

	if c.tail == c.head {
		delete(c.store, c.tail.key)
		c.tail = nil
		c.head = nil
		c.size--

		return
	}

	delete(c.store, c.tail.key)
	c.tail = c.tail.next
	c.size--
}

// Put puts variable value at variable key.
// Variable value is stored at the head of the cache.
// If the cache is full, variable at the tail of the cache is removed.
func (c *LRUCache) Put(key interface{}, value interface{}) {
	item, ok := c.store[key]

	if !ok {
		if c.size >= c.cap {
			c.evict()
		}

		newItem := &LinkedNode{
			key:   key,
			value: value,
			prev:  c.head,
		}

		c.promote(newItem)
		c.store[key] = newItem
		c.size++

		return
	}

	c.promote(item)
	item.value = value
}
