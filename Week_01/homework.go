// 641. 设计循环双端队列
// https://leetcode-cn.com/problems/design-circular-deque/

// 循环双端队列 CircularDoubleEndedQueue
type MyCircularDeque struct {
	data []int	// 数组(切片)存储数据
	// 实际容量为data容量减1，可以在这里记录也可以不记录，我选择不记录
	front int	// “头指针”数组下标
	last int	// “尾指针”数组下标
}

//    MyCircularDeque(k)：构造函数,双端队列的大小为k。
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		// 构造长度容量都为k+1的切片；当然也可以初始化为长度为0，容量为k+1；
		data: make([]int, k+1, k+1),	// 初始值全0
		front:0,
		last:0,
	}
}

//    insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
func (this *MyCircularDeque) InsertFront(value int) bool {
	// 检查队列是否已满
	if this.IsFull() {
		return false
	}

	// 插入元素。
	this.front = (len(this.data) + this.front - 1) % len(this.data)	// 先循环左移一位
	this.data[this.front] = value	// 填入数据

	return true
}

//    insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
func (this *MyCircularDeque) InsertLast(value int) bool {
	// 检查队列是否已满
	if this.IsFull() {
		return false
	}

	// 插入元素
	this.data[this.last] = value	// 填入数据
	this.last = (this.last + 1) % len(this.data)	// 循环右移一位

	return true
}

//    deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
func (this *MyCircularDeque) DeleteFront() bool {
	// 检查队列是否为空
	if this.IsEmpty() {
		return false
	}

	// 删除头部元素
	//this.data[this.front] = 0	// 置0，这一部完全不是必须，只是为了方便输出调试。可以将这句直接注释
	this.front = (this.front + 1) % len(this.data)	// 循环右移一位

	return true
}

//    deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
func (this *MyCircularDeque) DeleteLast() bool {
	// 检查队列是否为空
	if this.IsEmpty() {
		return false
	}

	// 删除头部元素
	this.last = (len(this.data) + this.last - 1) % len(this.data)	// 循环左移一位
	//this.data[this.last] = 0	// 置0，这一部完全不是必须，只是为了方便输出调试。可以将这句直接注释

	return true
}

//    getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
func (this *MyCircularDeque) GetFront() int {
	// 检查队列是否为空
	if this.IsEmpty() {
		return -1
	}

	// 获取头部元素
	return this.data[this.front]
}

//    getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
func (this *MyCircularDeque) GetRear() int {
	// 检查队列是否为空
	if this.IsEmpty() {
		return -1
	}

	// 获取尾部元素。这里要注意下，last应该循环左移一位得到数据下标
	return this.data[(len(this.data) + this.last - 1) % len(this.data)]
}

//    isEmpty()：检查双端队列是否为空。
func (this *MyCircularDeque) IsEmpty() bool {
	return this.last == this.front
}

//    isFull()：检查双端队列是否满了。
func (this *MyCircularDeque) IsFull() bool {
	return (this.last + 1) % len(this.data) == this.front
}


// 接雨水
// https://leetcode-cn.com/problems/trapping-rain-water/
func trap(height []int) int {
    var left,right,result,leftMax,rightMax int
    left = 0
    right = len(height)-1
    for left < right{
        if height[left] < height[right]{
            if height[left] > leftMax{
                leftMax = height[left]
            } else {
                result += leftMax - height[left]
            }
            left++
        } else {
            if height[right] > rightMax{
                rightMax = height[right]
            } else {
                result += rightMax - height[right]
            }
            right--
        }
    }
    return result
}