package main

/**
  map的key是可以比较的，通常是int和string,如果是struct，则struct不可变
  解决map并发的办法：加读写锁、分片降低锁的粒度，
  rwmutext加map 与 sync.map对比
  1.在读多写少的情况下适用sync.map,追加写，不适用于更新和写大于读的场景
  2.多个goroutine之间不想交的键的读和写适用于sync.map

  并发场景下这三种map使用做个对比比较好！！！
  sync.Map 、 map + rwmutex 和 第三方的currentmap(基于分段锁)
*/

func main() {

}
