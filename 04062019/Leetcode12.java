package com.jiang.problems;

public class Leetcode12 {
    public static String intToRoman(int num) {
    	String result="";
        int[] numbers= {1000,500,100,50,10,5,1};
        String[] Roman= {"M","D","C","L","X","V","I"};
        for(int i=0;i<numbers.length;i++) {
        	int wei=num/numbers[i];
        	if(wei==1&&(i==1||i==3||i==5)&&(num/9/numbers[i+1]==1)) {
        		result=result+Roman[i+1]+Roman[i-1];
            	num-=9*numbers[i+1];
        	}else if(wei==4) {
        		result=result+Roman[i]+Roman[i-1];
            	num-=wei*numbers[i];
        	}else {
            	num-=wei*numbers[i];
            	for(int j=wei;j>0;j--) {
            		result+=Roman[i];
            	}
        	}
        }
    	return result;
    }
    public static void main(String[] args) {
		System.out.println(intToRoman(1994));
	}
}
