def factorial (n: Int) : Int = {
  def loop(n: Int, acc: Int): Int = 
    if (n <= 0) acc
    else loop(n - 1, n * acc)
  loop(n, 1)
}
