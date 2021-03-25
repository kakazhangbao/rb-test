package main

import "fmt"

const (
	RED   int = 0
	BLACK int = 1
)

type RBNode struct {
	Co     int
	Left   *RBNode
	Right  *RBNode
	Parent *RBNode
	Data   int
}

type RBTree struct {
	Root *RBNode
}

func NewNode(data int) *RBNode {
	return &RBNode{
		Co:     RED,
		Left:   nil,
		Right:  nil,
		Parent: nil,
		Data:   data,
	}
}

func NewRBtree(rootData int) *RBTree {
	rootNode := NewNode(rootData)
	rootNode.Co = BLACK
	return &RBTree{
		Root: rootNode,
	}
}

func (t *RBTree)InsertData(data int) {
	//先插入到树中
	node := NewNode(data)
	t.insertNode(t.Root,node)
	t.insertFix(node)
}

func (t *RBTree)insertNode(root *RBNode,newNode *RBNode ){
	if root == nil {
		return
	}
	if newNode.Data < root.Data {
		if root.Left == nil {
			root.Left = newNode
			newNode.Parent = root
		} else{
			t.insertNode(root.Left,newNode)
		}
	}else{
		if root.Right == nil{
			root.Right = newNode
			newNode.Parent = root
		}else{
			t.insertNode(root.Right,newNode)
		}
	}
}

//获取叔叔节点
func (t *RBTree)GetUncle(z *RBNode ) *RBNode{
	if z.Parent == nil {
		return nil
	}
	if z.Parent.Parent ==nil {
		return nil
	}
	if z.Parent == z.Parent.Parent.Left {
		return z.Parent.Parent.Right
	}else {
		return z.Parent.Parent.Left
	}
}

func (t *RBTree)insertFix(z *RBNode ){
	//场景一  平衡节点为根节点
	//只需要变色即可
	if z.Parent == nil{
		z.Co = BLACK
		return
	}
	//场景二 平衡节点父节点为黑色
	// 平衡完成返回
	if z.Parent.Co == BLACK {
		return
	}
	//场景三 父节点为红色，叔叔节点为红色(到达此步，该节点父节点为红色，必定有祖父节点)
	//父节点和叔叔节点变黑色，祖父节点变为红色，递归处理
	UN :=t.GetUncle(z)
	GP := z.Parent.Parent
	P := z.Parent
	if UN != nil && UN.Co == RED {
		GP.Co = RED
		P.Co=BLACK
		UN.Co = BLACK
		t.insertFix(GP)
	}
	//到达场景四 说明必须要旋转了，如果父节点和当前节点在同一边，旋转一次，不在同一边需要旋转两次
	//这里注意由于所有指向都是双向的
	if z.Parent.Left == z && GP.Left == z.Parent { //4.1 当前节点和父节点都是左子叶
		//右旋节点
		GPP := GP.Parent
		if GPP != nil {
			if GPP.Left == GP {
				GPP.Left = P
			}else {
				GPP.Right = P
			}
			P.Parent = GPP
		}else{ //替换根节点
			t.Root = P
			P.Parent = nil
		}
		GP.Left = P.Right
		if P.Right != nil {
			P.Right.Parent = GP
		}
		GP.Co =RED
		P.Right = GP
		GP.Parent = P
		P.Co = BLACK
		return
	} else if z.Parent.Right == z && GP.Right == z.Parent {  //4.2 当前节点和父节点都是右子叶
		//左旋
		//判断是否是根节点
		GPP := GP.Parent
		if GPP != nil {
			if GPP.Left == GP {
				GPP.Left = z.Parent
			}else {
				GPP.Right = z.Parent
			}
			P.Parent = GPP
		}else{
			t.Root = P
			P.Parent = nil
		}
		GP.Right = P.Left
		if P.Left != nil {
			P.Left.Parent = GP
		}
		GP.Co = RED
		P.Left = GP
		GP.Parent = P
		P.Co = BLACK
		return

	}else if z.Parent.Left == z && GP.Right == z.Parent { //父节点为右子叶 和 当前节点为左子叶，需要两次旋转
		GP.Right = z
		z.Parent = GP

		P.Left = z.Right
		if z.Right != nil {
			z.Right.Parent = P
		}

		z.Right = P
		P.Parent = z
		t.insertFix(P)
	}else { //父节点为左子叶，当前节点为右子叶 需要两次旋转
		GP.Left = z
		z.Parent = GP


		P.Right = z.Left
		if z.Left != nil {
			z.Left.Parent = P
		}

		z.Left = P
		P.Parent = z
		//当前修复节点为父节点
		t.insertFix(P)
	}
}

//中序遍历
func PrintTree(rb *RBNode){
	if rb.Left != nil{
		PrintTree(rb.Left)
	}
	fmt.Printf( "%d ",rb.Data)
	if rb.Right != nil {
		PrintTree(rb.Right)
	}
}

//删除一个值
func (t *RBTree)DeleteData(data int){
	node := t.SearchData(t.Root,data)
	if node != nil{
		t.DeleteNode(node)
	}
}
//查找一个值
func (t *RBTree)SearchData(node *RBNode,data int) *RBNode{
	if node == nil{
		return node
	}
	if node.Data == data {
		return node
	} else if node.Data > data{
		return t.SearchData(node.Left,data)
	} else{
		return t.SearchData(node.Right,data)
	}
}

//删除一个节点
func (t *RBTree)DeleteNode(n *RBNode){
	//先找到继任节点
	var j *RBNode  //继任节点
	if n.Left!=nil && n.Right!=nil{
		j = t.findMin(n.Right)
		n.Data = j.Data
		t.DeleteNode(j)
		return
	}else if n.Left==nil && n.Right == nil{
		j = nil
	}else if n.Right!=nil{
		j= n.Right
	}else{
		j= n.Left
	}

	p := n.Parent
	if p != nil {
		if p.Left == n { //当前节点是左节点
			p.Left = j
		}else{//当前节点是右节点
			p.Right = j
		}
		if j != nil {
			j.Parent = p
		}
	}else{
		t.Root = j
	}
}

//查找一颗树的最小节点
func(t *RBTree)findMin(n *RBNode)*RBNode{
	if n == nil {
		return n
	}
	for {
		if n.Left == nil {
			break
		}
		n = n.Left
	}
	return n
}

//删除后修正红黑树
func DeleteFix(){
	//todo
}


func main(){
	//测试程序
	rb_tree := NewRBtree(100)
	rb_tree.InsertData(10)
	rb_tree.InsertData(20)
	rb_tree.InsertData(150)
	rb_tree.InsertData(300)
	rb_tree.InsertData(117)
	rb_tree.InsertData(1)
	rb_tree.InsertData(222)
	rb_tree.InsertData(223)
	rb_tree.InsertData(223)
	rb_tree.InsertData(2)
	rb_tree.InsertData(40)
	rb_tree.InsertData(777)
	rb_tree.InsertData(56)
	rb_tree.InsertData(26)
	PrintTree(rb_tree.Root)
}
