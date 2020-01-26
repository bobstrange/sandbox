# app

## Setup

```shell
npm install -D \
  typescript
```

`tsconfig.json`

```json
{
  "compilerOptions": {
    "target": "ES2020",
    "outDir": "./dist",
    "rootDir": "./src"
  }
}
```

### Setup webpack

```shell
npm install -D \
  webpack \
  webpack-cli \
  ts-loader \
  webpack-dev-server
```

webpack.config.js

```js
module.exports = {
  mode: "development",
  entry: "./src/index.ts",
  output: { filename: "bundle.js" },
  resolve: { extensions: [".ts", ".js"] },
  module: {
    rules: [
      { test: /\.ts/, use: "ts-loader", exclude: /node_modules/ }
    ]
  }
}
```

Run webpack `npx webpack`.
Run code `node dist/bundle.js`

Add webpack-dev-server configuration.

webpack.config.js

```js

  // ...
  devServer: {
    contentBase: "./assets",
    port: 4000
  }
}
```
