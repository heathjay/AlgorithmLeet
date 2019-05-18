package com.jiang.problems;

public class printmatrix {
	public static void printMatrixZigZag(int [][] matrix) {
		int tR=0;                //���Ͻ����꣨tR��tC�����½����꣨dR��dC��                   
		int tC=0;                //���Ͻ����������ž����һ���ƶ�tC++�������һ�����ұ�Ԫ�غ������ž������һ����
		int dR=0;                //���½����������ž����һ���ƶ�dC++�������һ�����±�Ԫ�غ������ž������һ����
		int dC=0;
		int endR=matrix.length-1;
		int endC=matrix[0].length-1;
		boolean fromUp=false;
		while(tR!= endR+1) {                                  //����������������깲ͬ�ƶ�����ӡб���ϵ�Ԫ�ؼ���
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
