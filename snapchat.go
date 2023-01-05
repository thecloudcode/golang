package main
import ("fmt")

func main(){
	// your code goes here
	var t int
	fmt.Scanln(&t)
	for i:=0;i<t;i++{
	    var n int
	    fmt.Scanln(&n)
	    var ar=make([]float64,n)
	    for i:=0;i<n;i++{
	        fmt.Scanf("%v",&ar[i])
	    }
	    var arr=make([]float64,n)
	    for i:=0;i<n;i++{
	        fmt.Scanf("%v",&arr[i])
	    }
	    ar=append(ar,0)
	    arr=append(arr,0)
	    m:=0
	    mm:=0
	    for i:=0;i<n+1;i++{
	        if ar[i]>0 && arr[i]>0{
	           // fmt.Println(ar[i],arr[i])
	            m++
	        }else{
	            if m>mm{
	                mm=m
	            }
	            m=0
	        }
	    }
	    fmt.Println(mm)
	   // var x=make([]float64,n)
	   // var c float64
	   // for i:=0;i<n;i++{
	   //     if ar[i]>0 && arr[i]>0{
	   //         c+=1.00
	   //     }else{
	   //         x=append(x,c)
	   //         c=0
	   //     }
	   // }
	   // fmt.Println(math.Max(x))
	}
}
