package com.jiang.problems;

import java.util.HashMap;

public class leetcode1 {
    public static int[] twoSum(int[] nums, int target) {
        HashMap<Integer,Integer> find=new HashMap<Integer,Integer>();
        for(int i=0;i<nums.length;i++) {
        	if(find!=null&&find.get(nums[i])!=null) {
        		return new int[] {find.get(nums[i]),i};
        	}
        	find.put(target-nums[i], i);
        }
        return null;
    }
    public static void main(String[] args) {
    	int[] nums= {2,7,11,5};
		int[] shuzu=twoSum(nums,9);
		for(int i=0;i<shuzu.length;i++) {
			System.out.print(shuzu[i]+",");
		}
		
	}
}
