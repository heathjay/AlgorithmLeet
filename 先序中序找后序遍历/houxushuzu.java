package com.jiang.problems;

import java.util.HashMap;

public class houxushuzu {
	public int[] getPosArray(int[] pre,int[] in) {
		if(pre==null||in==null) {
			return null;
		}
		int len=pre.length;
		int[] pos=new int[len];
		HashMap<Integer,Integer> map=new HashMap<Integer,Integer>();
		for(int i=0;i<len;i++) {
			map.put(in[i], i);//中序数组对应存储其值和位置
		}
		setPos(pre,0,len-1,in,0,len-1,pos,len-1,map);//pre先序遍历，后面是它的开始和结束位置，in中序遍历，pos后序遍历的数组
		return pos;
	}
	//整个函数思想是把先序遍历的第一个值找到，即头结点这样就能将它直接放到我们要返回的数组的最后一个
	//然后通过中序遍历找到左右子树再递归找到左右子树的头结点放到后序节点相应的位置上面
	private int setPos(int[] p, int pi, int pj, int[] n, int ni, int nj, int[] s, int si,
			HashMap<Integer, Integer> map) {
			if(pi>pj) {
				return si;
			}
			s[si--]=p[pi];
			int i=map.get(p[pi]);
			si=setPos(p,pj-nj+i+1,pj,n,i+1,nj,s,si,map);//头结点的右子树找头结点
			return setPos(p,pi+1,pi+i-ni,n,ni,i-1,s,si,map);//头结点的左子树找它的头结点
	}
}
