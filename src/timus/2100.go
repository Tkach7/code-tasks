package main
import (
    "fmt"
    "strings"
)

func main() {
    const Cost_one_person = 100
    var guestCount, price int
    var guest string
    fmt.Scan(&guestCount)

    for i := 0; i < guestCount; i++ {
        fmt.Scan(&guest)
        price += len(strings.Split(guest, "+")) * Cost_one_person
    }

    price = price + 200

    if price == 1300 {
       price = price + 100
    }


    fmt.Println(price)
}

