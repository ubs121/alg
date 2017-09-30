package main
import ( "bufio" ; "fmt" ; "os" )

var array [200]int

/******************************************************    
 * QuickSort -  Массивын өгөгдсөн мужийг эрэмбэлэх    *
 *                                                    *
 * Параметрүүд                                        *
 *   low –- эрэмбэлэх мужийн доод хязгаар.            *
 *   high –- эрэмбэлэх мужийн дээд хязгаар.           *
 *                                                    *
 ******************************************************/
func QuickSort(low int, high int) {
    i:= low; j:= high
    pivot:=array[(low + high)/2]
    var temp int

    for i<=j {
        for array[i] < pivot { i++ }
        for array[j] > pivot { j-- }
        if i <= j {
            temp = array[i]
            array[i] = array[j]
            array[j] = temp
            i++; j--
        }
    }

    if low < j { QuickSort(low, j) }    /* зүүн хэсгийг эрэмбэлэх */
    if i < high { QuickSort(i, high) } /* баруун хэсгийг эрэмбэлэх */
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    s, _ := reader.ReadString('\n')

    /* элементийн тоо */
    n:=0

    /* Сул зайгаар тусгаарлагдсан тоон дарааллыг унших */
    for i:=0; i<len(s); n++{
        for i<len(s) && s[i]==' ' { i++ }
        fmt.Sscanf(s[i:], "%d", &array[n])
        for i<len(s) && s[i]!=' ' { i++ }
    }

    /* массивыг эрэмбэлэх */
    QuickSort(0, n - 1)

    /* эрэмбэлэгдсэн  массивыг хэвлэж харуулах */
    for i:= 0; i<n; i++ {
       fmt.Printf("%d ", array[i])
    }
}