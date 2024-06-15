package main
import "fmt"

type cataloge interface{
      shipping()float64
      tax()float64
}
type confi struct{
  name string
  price , qty float64
}

func (c* confi)tax()float64{
return c.price * c.qty * 0.05
}
func (c*confi)shipping() float64{
return c.qty * 5
}

func main(){
  a:=confi{}
  a.price=500
  a.qty=2
  fmt.Println("shipping Charge:",a.shipping())
  fmt.Println("Tax:",a.tax())

}
