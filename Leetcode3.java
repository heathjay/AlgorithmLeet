package com.jiang.problems;

import java.util.HashMap;

public class Leetcode3 {
    public static int lengthOfLongestSubstring(String s) {
    	if(s.equals(" ")) {
    		return 1;
    	}
        if(s.equals("")||s==null){
            return 0;
        }
        int length=0;
        int once=0;
        int pre=-1;
        HashMap<String,Integer> index=new HashMap<String,Integer>();
        for(int i=0;i<s.length();i++) {
        	String sample=s.substring(i, i+1);
        	if(index.get(sample)!=null) {
        	pre=Math.max(pre, index.get(sample));
        	}
        	once=i-pre;
        	length=Math.max(once, length);
        	index.put(sample, i);
        }
		length=Math.max(length, once);
        return length;
    }
    public static void main(String[] args) {
    	String s="bybz";
    	int big=lengthOfLongestSubstring(s);
    	System.out.println(big);
	}
}
