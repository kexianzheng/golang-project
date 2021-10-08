
package main

import "fmt"

func main(){
     /*chengfakoujuebiao
     */
    /* for j :=1;j<=1;j++{
          fmt.Printf("%d * %d = %d\t", 1, j, 1*j)
     }
     fmt.Println()
     for j :=1;j<=2;j++{
          fmt.Printf("%d * %d = %d\t", 2, j, 2*j)
     }
     fmt.Println()
     for j :=1;j<=3;j++{
          fmt.Printf("%d * %d = %d\t", 3, j, 3*j)
     }
     fmt.Println()
     for j :=1;j<=4;j++{
          fmt.Printf("%d * %d = %d\t", 4, j, 4*j)
     }
     fmt.Println()*/
     for i := 1;i<10; i++{
          for j :=1;j<=i;j++{
               fmt.Printf("%d * %d = %d\t", j, i, i*j)
          }
          fmt.Println()
     }
}