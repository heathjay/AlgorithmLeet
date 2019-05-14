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
			map.put(in[i], i);//���������Ӧ�洢��ֵ��λ��
		}
		setPos(pre,0,len-1,in,0,len-1,pos,len-1,map);//pre������������������Ŀ�ʼ�ͽ���λ�ã�in���������pos�������������
		return pos;
	}
	//��������˼���ǰ���������ĵ�һ��ֵ�ҵ�����ͷ����������ܽ���ֱ�ӷŵ�����Ҫ���ص���������һ��
	//Ȼ��ͨ����������ҵ����������ٵݹ��ҵ�����������ͷ���ŵ�����ڵ���Ӧ��λ������
	private int setPos(int[] p, int pi, int pj, int[] n, int ni, int nj, int[] s, int si,
			HashMap<Integer, Integer> map) {
			if(pi>pj) {
				return si;
			}
			s[si--]=p[pi];
			int i=map.get(p[pi]);
			si=setPos(p,pj-nj+i+1,pj,n,i+1,nj,s,si,map);//ͷ������������ͷ���
			return setPos(p,pi+1,pi+i-ni,n,ni,i-1,s,si,map);//ͷ����������������ͷ���
	}
}
