package com.jiang.problems;

import javax.swing.plaf.synth.SynthScrollBarUI;

public class Levenshtein {
	public static int[][] getmatric(String a,String b){
		int[][] i=new int [a.length()+1][b.length()+1];
		for(int j=0;j<a.length()+1;j++) {
			i[j][0]=j;
		}
		for(int j=0;j<b.length()+1;j++) {
			i[0][j]=j;
		}
		for(int al=1;al<a.length()+1;al++) {
			for(int bl=1;bl<b.length()+1;bl++) {
					i[al][bl]=Math.min(a.charAt(al-1)==b.charAt(bl-1)?i[al-1][bl-1]:i[al-1][bl-1]+1, Math.min(i[al][bl-1]+1, i[al-1][bl]+1));
			}
		}
		return i;
	}
	public static void main(String args[]) {
		int[][] i=getmatric("batyu","beauty");
		for(int j=0;j<i.length;j++) {
			for(int k=0;k<i[0].length;k++) {
				System.out.print(i[j][k]+" ");
			}
			System.out.println();
		}
	}
}
