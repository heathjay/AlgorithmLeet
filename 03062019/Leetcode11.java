package com.jiang.problems;

public class Leetcode11 {
    public int maxArea(int[] height) {
    	int max=0;
        for(int i=0;i<height.length;i++) {
        	for(int j=i;j<height.length;j++) {
        		int mianji=(j-i)*Math.min(height[i], height[j]);
        		max=Math.max(max, mianji);
        	}
        }
    	return max;
    }
}
