package mysort

func BubbleSort(arr []int)  {
	l:= len(arr)
	for i:= l-1; i >= 0; i-- {
		for j:= 0; j<= i-1 ;j++ {
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
}

func SelectSort(arr []int){
	l:= len(arr)
	for i:= l-1; i >= 0; i-- {
		max := i
		for j:=0;j<i;j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}
		if max != i {
			arr[i],arr[max] = arr[max] ,arr[i]
		}
	}
}

func HeapSort(arr []int){
	l:= len(arr)
	for i := l/2 ;i>= 0;i-- {
		HeapFix(arr,i,l)
	}
	for j:= l-1; j>0; j-- {
		arr[0],arr[j] = arr[j] ,arr[0]
		l-=1
		HeapFix(arr,0,l)
	}
}

func HeapFix(arr []int ,index ,length int){
	lson := index*2 +1
	rson := index*2 + 2
	maxindex := index
	if lson < length && arr[lson] > arr[maxindex] {
		maxindex = lson
	}
	if rson< length && arr[rson] > arr[maxindex] {
		maxindex = rson
	}
	if maxindex != index {
		arr[maxindex],arr[index] = arr[index], arr[maxindex]
		HeapFix(arr,maxindex,length)
	}
}

func InsertSort(arr []int){
	l:= len(arr)
	for i:=1 ;i<= l-1;i++{
		tmp := arr[i]
		j:=0
		for j= i;j > 0 && arr[j-1] > tmp ;j--{
			arr[j] =arr[j-1]
		}
		arr[j] = tmp
	}
}


func ShellSort(arr []int){
	l:= len(arr)
	for d := l/2 ; d>0 ;d/=2 {
		for i:= d ; i < l ; i++ {
			for j:= i-d ;j>=0; j-=d {
				if arr[j] > arr[j+d] {
					arr[j] ,arr[j+d] = arr[j+d] ,arr[j]
				}
			}
		}
	}
}

func FastSort(arr []int){
	l := len(arr)
	if l < 1 {
		return
	}
	qviod := arr[0]
	lok := false
	lindex := 0
	rindex := l-1
	for lindex != rindex {
		if lok{
			if arr[lindex] > qviod {
				arr[rindex] = arr[lindex]
				lok = false
				rindex--
			}else{
				lindex++
				lok= true
			}
		}else{
			if arr[rindex] < qviod{
				arr[lindex] =arr[rindex]
				lindex++
				lok =true
			}else{
				rindex--
				lok =false
			}
		}
	}
	arr[rindex] =qviod
	if lindex > 1 {
		FastSort(arr[0:lindex])
	}
	if l-lindex > 2  {
		FastSort(arr[lindex+1:])
	}
}

func MergeSort(arr []int) []int{
	l:= len(arr)
	if l < 2 {
		return arr
	}
	mid := l/2
	return Merge(MergeSort(arr[:mid]),MergeSort(arr[mid:]))
}

func Merge(arr []int , arr2 []int) []int{
	res := make([]int,0)
	for len(arr) >0 && len(arr2) >0 {
		if arr[0] < arr2[0] {
			res = append(res,arr[0])
			arr = arr[0:]
		}else{
			res =append(res,arr2[0])
			arr2 = arr2[0:]
		}
	}
	if len(arr) > 0 {
		res = append(res,arr...)
	}else{
		res = append(res,arr2...)
	}
	return res
}
