package main

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	cache    map[int]*Node
	head     *Node // most recently used
	tail     *Node // least recently used
	capacity int
	size     int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		cache:    make(map[int]*Node),
		head:     nil, // start with empty list
		tail:     nil,
		capacity: capacity,
		size:     0,
	}
}

func (c *LRUCache) Get(key int) int {
	if node, exists := c.cache[key]; exists {
		c.moveToFront(node)
		return node.value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, exists := c.cache[key]; exists {
		// Key already exists, update value and move to front, no need to update size
		node.value = value
		c.moveToFront(node)
	} else {
		// New key, create new node
		node := &Node{key: key, value: value}
		c.cache[key] = node
		c.addToFront(node)
		c.size++

		// If we exceeded capacity, remove the least recently used (tail)
		if c.size > c.capacity {
			removedNode := c.removeFromBack()
			delete(c.cache, removedNode.key)
			c.size--
		}
	}
}

func (c *LRUCache) moveToFront(node *Node) {
	c.removeFromList(node)
	c.addToFront(node)
}

func (c *LRUCache) removeFromList(node *Node) {
	// Handle case where this is the only node
	if c.head == node && c.tail == node {
		c.head = nil
		c.tail = nil
		return
	}

	// Handle case where this is the head node
	if c.head == node {
		c.head = node.next
		if c.head != nil {
			c.head.prev = nil
		}
		return
	}

	// Handle case where this is the tail node
	if c.tail == node {
		c.tail = node.prev
		if c.tail != nil {
			c.tail.next = nil
		}
		return
	}

	// Handle case where this is a middle node
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) addToFront(node *Node) {
=	node.prev = nil
	node.next = nil

	// If list is empty, this becomes both head and tail
	if c.head == nil {
		c.head = node
		c.tail = node
		return
	}

	// Add to front of existing list
	node.next = c.head
	c.head.prev = node
	c.head = node
}

func (c *LRUCache) removeFromBack() *Node {
	if c.tail == nil {
		return nil
	}

	nodeToRemove := c.tail
	c.removeFromList(nodeToRemove)
	return nodeToRemove
}
