package com.jiang.problems;

import java.util.Stack;

public class Node {
	public int value;
	public Node left;
	public Node right;
	public Node(int data) {
		this.value=data;
	}
	//递归先序遍历
	public void preOrderRecur(Node head) {
		if(head==null) {
			return;
		}
		System.out.print(head.value+"");
		preOrderRecur(head.left);
		preOrderRecur(head.right);
	}
	//递归中序遍历
	public void inOrderRecur(Node head) {
		if(head==null) {
			return;
		}
		inOrderRecur(head.left);
		System.out.print(head.value+"");
		inOrderRecur(head.right);
	}
	//递归后序遍历
	public void posOrderRecur(Node head) {
		if(head==null) {
			return;
		}
		posOrderRecur(head.left);
		posOrderRecur(head.right);
		System.out.print(head.value+"");
	}
	//非递归先序遍历
	//通过压栈，获取头结点后按顺序压入右子树再左子树，然后每次获取栈头节点并压入其左右节点得到结果
	public void preOrderUnRecur(Node head) {
		if(head!=null) {
			Stack<Node> stack=new Stack<Node>();
			stack.add(head);
			while(!stack.isEmpty()) {
				head=stack.pop();
				System.out.print(head.right);
				if(head.right!=null) {
					stack.push(head.right);
				}
				if(head.left!=null) {
					stack.push(head.left);
				}
			}
		}
	}
	//非递归中序遍历需要先遍历左边再头结点再右边，就可以先把栈放入所有左边的节点然后弹出一个如果
	//该节点没有左子树的时候就直接打印并遍历其右子树，该节点有左子树就继续压栈其左子树
	public void inOrderUnRecur(Node head) {
		if(head!=null) {
			Stack<Node> stack=new Stack<>();
			while(!stack.isEmpty()||head!=null) {
				if(head!=null) {
					stack.push(head);
					head=head.left;
				}else {
					head=stack.pop();
					System.out.print(head.value+" ");
					head=head.right;
				}
			}
		}
	}
	//后序遍历使用两个栈，第一个先弹入头节点并弹到另外一个栈中然后弹入左节点再弹入右节点
	//全部遍历完以后将这个栈弹出把全部东西弹入到另一个栈中，这样另外一个栈的顺序就变成了左、右、中了
	public void posOrderUnRecurl(Node head) {
		if(head!=null) {
			Stack<Node> olds=new Stack<Node>();
			Stack<Node> news=new Stack<Node>();
			olds.push(head);
			while(!olds.isEmpty()) {
				head=olds.pop();
				news.push(head);
				if(head.left!=null) {
					olds.push(head.left);
				}
				if(head.right!=null) {
					olds.push(head.right);
				}
			}
			while(!news.isEmpty()) {
				System.out.print(news.pop().value+" ");
			}
		}
	}
}
