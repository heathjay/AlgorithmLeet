package com.jiang.problems;

public class reversestring {
	public static void rotateWord(char[] chas) {
		if(chas==null||chas.length==0) {
			return;
		}
		reverse(chas,0,chas.length-1);//JCP is my son->nos ym si PCJ
		int l=-1;
		int r=-1;
		for(int i=0;i<chas.length;i++) {
			if(chas[i]!=' ') {
				l=i==0||chas[i-1]==' '?i:l;//�ַ���Ϊ��ͷʱ���߿ո��һ��ʱ��ֵl��ȡλ��
				r=i==chas.length-1||chas[i+1]==' '?i:r;//�ַ���Ϊ��βʱ������һ���ַ�Ϊ���ʽ�����λ�ø�ֵr
			}
			if(i!=-1&&r!=-1) {     //���ҵ����ʿ�ͷ�ͽ�βʱת��
				reverse(chas,l,r);//���ʵ���˳�� nos->son
				l=-1;
				r=-1;
			}
		}
	}

	public static void reverse(char[] chas, int start, int end) {
		char tmp=0;
		while(start<end) {
			tmp=chas[start];
			chas[start]=chas[end];
			chas[end]=tmp;
			start++;
			end--;
		}
		
	}
	public static void main(String[] args) {
		String temp="JSP is my son";
		char  chas[]=new char[temp.length()];
		for(int i=0;i<temp.length();i++) {
			chas[i]=temp.charAt(i);
		}
		rotateWord(chas);
		for(int i=0;i<temp.length();i++) {
			System.out.print(chas[i]);
		}
		
	}
}
