# CMake Tutorial のメモ

[Document](https://termoshtt.gitlab.io/cmake-book/first_cmake.html)

## getting started

まずは、cmake を動かしてみる [./getting_started](./getting_started/)

- CMakeLists.txt を用意
- ソースコードも用意
- `cmake` を実行
  - `cmake .`
  - 色々ファイルが出力される
    - cmake_install.cmake
    - CMakeFiles
    - CMakeCache.txt
    - Makefile
- Makefile が出来ているので `make` を実行
- CMakeLists.txt に定義されている実行ファイル (ここでは Main) が生成される
