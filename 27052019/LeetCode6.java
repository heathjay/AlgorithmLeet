package com.jiang.problems;

import java.util.HashMap;

public class LeetCode6 {
    public static String convert(String s, int numRows) {
    	int group=1;
    	String result="";
    	HashMap<Integer,String> hashmap=new HashMap<Integer,String>();
    	for(int i=0;i<numRows;i++) {
    		hashmap.put(i, "");
    	}
    	if(numRows==1) {
    		return s;
    	}
    	if(numRows==2) {
    		int[] order= {0,1,2};
        	for(int j=0;j<s.length();j++) {
        		String temp=s.substring(j,j+1);
        		temp=hashmap.get(order[j%numRows])+temp;
        		hashmap.put(order[j%numRows], temp);
        		}
        	for(int i=0;i<numRows;i++) {
        		result+=hashmap.get(i);
        	}
        	return result;
    	}
    	int order[]=new int[2*numRows-2];
    	for(int i=0;i<2*numRows-2;i++) {
    		if(i<numRows) {
        		order[i]=i;
    		}else {
    			order[i]=numRows-group-1;
    			group++;
    		}   		
    	}
    	for(int j=0;j<s.length();j++) {
    		String temp=s.substring(j,j+1);
    		temp=hashmap.get(order[j%(2*numRows-2)])+temp;
    		hashmap.put(order[j%(2*numRows-2)], temp);
    		}
    	for(int i=0;i<numRows;i++) {
    		result+=hashmap.get(i);
    	}
        return result;
    }
    public static void main(String[] args) {
		String result=convert("PAYPALISHIRING",4);
		System.out.print(result);
	}
}
