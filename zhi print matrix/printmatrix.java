package com.jiang.problems;

public class printmatrix {
	public static void printMatrixZigZag(int [][] matrix) {
		int tR=0;                //左上角坐标（tR，tC）左下角坐标（dR，dC）                   
		int tC=0;                //左上角坐标先沿着矩阵第一行移动tC++当到达第一行最右边元素后再沿着矩阵最后一列走
		int dR=0;                //左下角坐标先沿着矩阵第一列移动dC++当到达第一列最下边元素后再沿着矩阵最后一行走
		int dC=0;
		int endR=matrix.length-1;
		int endC=matrix[0].length-1;
		boolean fromUp=false;
		while(tR!= endR+1) {                                  //左上坐标和左下坐标共同移动，打印斜线上的元素即可
			printLevel(matrix,tR,tC,dR,dC,fromUp);
			tR=tC==endC?tR+1:tR;
			tC=tC==endC?tC:tC+1;
			dC=dR==endR?dC+1:dC;
			dR=dR==endR?dR:dR+1;
			fromUp=!fromUp;
		}
		System.out.println();
	}
	public static void printLevel(int[][]m,int tR,int tC,int dR,int dC,boolean f) {
		if(f) {
			while(tR!=dR+1) {
				System.out.print(m[tR++][tC--]+",");
			}
			}else {
				while(dR!=tR-1) {
					System.out.print(m[dR--][dC++]+",");
				}
			}
		}
	public static void main(String[] args) {
		int[][] m= {{1,2,3,4},{5,6,7,8},{9,10,11,12}};
		printMatrixZigZag(m);
	}
}
