//首先不光要寻找到左子树，还要在左子树中找到最大值，同理右子树也得找到最小值，在递归到头结点的时候保证头结点大于左子树最大值，小于右子树最小值。
//这时还要获取左子树最大搜索子树是否就是左子树的头结点，右子树同理，这样如果都是的话就说明最大搜索二叉树还能增大
//所以仅定义一棵树的头结点不够，还要封装结点的信息
public class ReturnTupe{
  public Node maxBSTHead;//最大搜索二叉树头结点
  public Node maxBSTSize;//最大的大小
  public int min;//子树最小值
  public int max;//子树最大值
}
//主方法
public Node getMaxBST(Node head){
  return prosess(head).maxBSTHead;
}
//递归方法
public ReturnType prosess(Node X){
if( X==null){
  return new ReturnType(null,0,Integer.MAX_VALUE,Integer.MIN_VALUE);
  }
  ReturnType lData=prosess(X.left);//得到左子树所有信息
  ReturnType rData=prosess(X.left);//得到右子树所有信息
  int min=Math.min(X.value,Math.min(lData.min,rData.min));//得到X的最小值
  int max=Math.min(X.value,Math.max(lData.max,rData.max));//得到X的最大值 
  int maxBSTSize=Math.max(lData.maxBSTSize,rData.maxBSTSize);//如果只考虑左右子树的最大搜索二叉树，不考虑把他们和当前节点合起来当成一个二叉树时的最大搜索二叉树值
  Node maxBSTHead=lData.maxBSTSize>=rData.maxBSTSize?lData.maxBSTHead:rData.maxBSTHead;//同理
  if(lData.maxBSTHead==X.left&&rData.maxBSTHead==X.right&&X.value>lData.max&&X.value<rData.min){//如果当前结点也可以加入到搜索二叉树中
    maxBSTSize=lData.maxBSTSize+rData.maxBSTSize+1;
    maxBSTHead=X;
  }
  return new ReturnType(maxBSTHead,maxBSTSize,min,max);
}
