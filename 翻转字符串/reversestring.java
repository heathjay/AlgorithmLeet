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
				l=i==0||chas[i-1]==' '?i:l;//字符串为开头时或者空格第一个时赋值l获取位置
				r=i==chas.length-1||chas[i+1]==' '?i:r;//字符串为结尾时或者下一个字符为单词结束是位置赋值r
			}
			if(i!=-1&&r!=-1) {     //当找到单词开头和结尾时转换
				reverse(chas,l,r);//单词调换顺序 nos->son
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
