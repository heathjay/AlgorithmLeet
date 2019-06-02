package com.jiang.problems;

public class Leetcode10 {
	/*
	 * 检验输入字符是否符合规则
	 */
	public static boolean isValid(char[]s,char[]e) {
		for(int i=0;i<s.length;i++) {
			if(s[i]<'0'||s[i]>'9') {
				return false;
			}
		}
		for(int j=0;j<e.length;j++) {
			if(e[j]<'0'||e[j]>'9'||(e[j]=='*'&&e[j-1]=='*')) {
				return false;
			}
		}
	return true;
	}
	/*
	 * 主函数
	 */
	public static boolean isMatch(String str,String exp) {
		if(str==null||exp==null) {
			return false;
		}
		char[] s=str.toCharArray();
		char[] a=exp.toCharArray();
		return isValid(s,a)?process(s,a,0,0):false;
	}
	/*
	 * 递归判断
	 */
	public static boolean process(char[] s, char[] e, int si, int ej) {
		if(ej==e.length) {
			return si==s.length;
		}
		/*
		 * 当目前规则没有*
		 */
		if(ej+1==e.length||e[ej+1]!='*') {
			return si!=s.length && (s[si]==e[ej]||e[ej]=='.') && process(s,e,si+1,ej+1);
		}
		/*
		 * 目前规则为*
		 */
		while(si!=s.length && (s[si]==e[ej])||e[ej]=='.') {
			if(process(s,e,si,ej+2)) {      //如果出现很多能匹配的e[ei]*而且e[ei]仅匹配一部分就能全部匹配则返回成功
				return true;
			}
			si++;
		}
		return process(s,e,si,ej+2);
	}
}
