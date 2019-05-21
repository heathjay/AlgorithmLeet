package com.jiang.problems;

public class lianbiaoadd {
	/**
	 * Definition for singly-linked list.
	 * public class ListNode {
	 *     int val;
	 *     ListNode next;
	 *     ListNode(int x) { val = x; }
	 * }
	 */
	    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
	        int wei=0;
			ListNode last=null;
			ListNode first=null;
	        int value=0;
			while(l1!=null||l2!=null) {
	            int one=l1!=null?l1.val:0;
	            int two=l2!=null?l2.val:0;
	            value=one+two+wei;
	            ListNode node=new ListNode(value%10);
	            wei=value/10;
	            if(first==null){
	                first=node;
	            }else{
	                last.next=node;
	            }
	            last=node;
	            l1=l1!=null?l1.next:null;
	            l2=l2!=null?l2.next:null;
	        }
	        if(wei==1){
	            last.next=new ListNode(1);
	        }
	        return first;
	    }
}
