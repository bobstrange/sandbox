/** Refers: http://qiita.com/kt3k@github/items/df783ae54caa4eeca2ae */

module.exports = function(config) {
  config.set({

    /** brwoserifyを追加 */
    basePath: '',
    frameworks: ['jasmine', 'browserify'],
    /** 実装ファイルはrequireで取得できるので、
     *  テストファイルだけの指定でOK */
    files: [
      'src/*.js',
      'src/**/*.js',
      'test/*.spec.js',
      'test/**/*.spec.js'
    ],
    exclude: [],
    /* 全てのfileにbrowserifyがかかる必要があるため、filesで指定したパターンと
     * 同様のものを指定する */
    preprocessors: {
      'src/*.js': ['browserify'],
      'src/**/*.js': ['browserify'],
      'test/*.spec.js': ['browserify'],
      'test/**/*.spec.js': ['browserify']
    },
    browserify: {
      debug: true,
      transform: [
        require('browserify-istanbul')({
          instrumenter: require('isparta'),
          ignore:       ['**/test/**']
        }),
        'babelify'
      ]
    },
    /** coverageを追加 */
    reporters: ['progress', 'coverage'],
    port: 9876,
    colors: true,
    logLevel: config.LOG_INFO,
    autoWatch: false,
    browsers: ['PhantomJS'],
    singleRun: true,
    concurrency: Infinity
  });
};
