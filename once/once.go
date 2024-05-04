package main

/**
  once用来执行且仅仅执行一次动作，用于单例对象的初始化场景
  sync.Once只暴露了一个方法do,你可以多次调用do方法，但是只有调用一次do方法时f参数才会执行，f参数无参无返回值
  sync.Once底层使用mutex和双检查机制保证f函数参数只执行了一次
  如果once.Do()初始化失败了，once还是会认为初始化成功，你不会再次执行f函数方法
*/

func main() {

}
