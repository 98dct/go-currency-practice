package main

/**
  sync.Pool 数据类型用来保存一组可独立访问的临时对象。请注意这里加粗的“临时”这两个字，
  它说明了 sync.Pool 这个数据类型的特点，也就是说，它池化的对象会在未来的某个时候被毫无预兆地移除掉。
  而且，如果没有别的对象引用这个被移除的对象的话，这个被移除的对象就会被垃圾回收掉
  sync.Pool并发安全，使用后不可以复制  new get put
*/

func main() {

}
