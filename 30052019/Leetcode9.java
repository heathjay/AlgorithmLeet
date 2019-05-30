package com.jiang.problems;

import java.util.Stack;

public class Leetcode9 {
    public boolean isPalindrome(int x) {
        Stack<String> stack=new Stack<String>();
        String notify=String.valueOf(x);
        String backwards="";
        for(int i=0;i<notify.length();i++) {
        	String a=notify.substring(i, i+1);
        	stack.push(a);
        }
        for(int i=0;i<notify.length();i++) {
        	backwards+=stack.pop();
        }
        return notify.equals(backwards);
    }
}
