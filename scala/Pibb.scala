def pibb(n: Int) : Int = {
  def loop(n: Int) : Int = {
    if (n <= 2) 1
    else loop(n-1) + loop(n-2)
  }
  loop(n)
}
