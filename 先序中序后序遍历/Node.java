package com.jiang.problems;

import java.util.Stack;

public class Node {
	public int value;
	public Node left;
	public Node right;
	public Node(int data) {
		this.value=data;
	}
	//�ݹ��������
	public void preOrderRecur(Node head) {
		if(head==null) {
			return;
		}
		System.out.print(head.value+"");
		preOrderRecur(head.left);
		preOrderRecur(head.right);
	}
	//�ݹ��������
	public void inOrderRecur(Node head) {
		if(head==null) {
			return;
		}
		inOrderRecur(head.left);
		System.out.print(head.value+"");
		inOrderRecur(head.right);
	}
	//�ݹ�������
	public void posOrderRecur(Node head) {
		if(head==null) {
			return;
		}
		posOrderRecur(head.left);
		posOrderRecur(head.right);
		System.out.print(head.value+"");
	}
	//�ǵݹ��������
	//ͨ��ѹջ����ȡͷ����˳��ѹ������������������Ȼ��ÿ�λ�ȡջͷ�ڵ㲢ѹ�������ҽڵ�õ����
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
	//�ǵݹ����������Ҫ�ȱ��������ͷ������ұߣ��Ϳ����Ȱ�ջ����������ߵĽڵ�Ȼ�󵯳�һ�����
	//�ýڵ�û����������ʱ���ֱ�Ӵ�ӡ�����������������ýڵ����������ͼ���ѹջ��������
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
	//�������ʹ������ջ����һ���ȵ���ͷ�ڵ㲢��������һ��ջ��Ȼ������ڵ��ٵ����ҽڵ�
	//ȫ���������Ժ����ջ������ȫ���������뵽��һ��ջ�У���������һ��ջ��˳��ͱ�������ҡ�����
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
