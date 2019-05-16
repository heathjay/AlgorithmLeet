package main

type Node struct {
	Left  *Node
	Right *Node
	data  float64
}

func findErr(res *[]float64, err *[]float64) {

	pre := (*res)[0]
	for i := 0; i < len(*res); i++ {
		if pre > (*res)[i] && len(*err) == 0 {
			(*err)[0] = pre
		}
		if pre > (*res)[i] && len(*err) != 1 {
			(*err)[1] = (*res)[i]
		}
		pre = (*res)[i]
	}

}

//先中序遍历
//然后形成递增序列

func InOrder(this *Node, res *[]float64) {
	if this == nil {
		return
	}
	InOrder(this.Left, res)
	*res = append(*res, this.data)
	InOrder(this.Right, res)
}

func main() {

}
