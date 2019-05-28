package com.jiang.problems;

public class Leetcode7 {
	public static int reverse(int x) {
		if(x==0) {
			return 0;
		}
		boolean fude=false;
		String nums=String.valueOf(x);
		String reversenums="";
		for(int i=nums.length();i>0;i--) {
			String number=nums.substring(i-1, i);
			reversenums+=number;
		}
		if(reversenums.startsWith("0")) {
			reversenums=reversenums.substring(1);
		}
		if(reversenums.contains("-")) {
			reversenums=reversenums.replace("-", "");
			fude=true;
		}
		Long data=Long.parseLong(reversenums);
		if(fude) {
			data=0-data;
		}
		if(data>Integer.MAX_VALUE||data<Integer.MIN_VALUE) {
			return 0;
		}else {
			return data.intValue();
		}

	}
	public static void main(String[] args) {
		System.out.println(reverse(0));
	}
}
