# tools

`tsc-watch`
sourceファイルの変更を検知してビルドするだけでなく、`--onsuccess`
でビルド成功後に実行したい処理を記述できる

`eslint`

```shell
npm install -D \
  eslint \
  @typescript-eslint/parser \
  @typescript-eslint/eslint-plugin \
  prettier \
  eslint-config-prettier
```

`eslintrc.js`

```js
module.exports = {
  root: true,
  parser: '@typescript-eslint/parser',
  plugins: [
    '@typescript-eslint'
  ],
  extends: [
    'eslint:recommended',
    'plugin:@typescript-eslint/eslint-recommended',
    'plugin:@typescript-eslint/recommended',
    'prettier/@typescript-eslint'
  ],
  env: {
    browser: true,
    node: true
  },
  rules: {
    "no-console": false
  }
}
```
