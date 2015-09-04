#include <iostream>
#include "sqlist.h"
using namespace std;

void InsertSort(SqList &L){
	int i,j;
	for(i=2;i<L.length;++i)
		if (LT(L.elem[i].key,L.elem[i-1].key)){
		L.elem[0] = L.elem[i];
		L.elem[i] = L.elem[i-1];
		for(j=i-2; LT(L.elem[i-1].key,L.elem[j].key); --j)
			L.elem[j+1] = L.elem[j];
		L.elem[j+1] = L.elem[0];
		}
}//InsertSort插入排序算法
//未完成算法部分

main(){
	cout << "init\n" << endl;
}
