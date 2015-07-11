var gulp = require('gulp');
var babel = require('gulp-babel');

var ES6_PATH = './es6/';
var COMPILED_FILE_PATH = './source/';

gulp.task('babel', function() {
  gulp.src(ES6_PATH + '*.es6')
    .pipe(babel())
    .pipe(gulp.dest(COMPILED_FILE_PATH));
});

gulp.task('watch', function() {
  gulp.watch(ES6_PATH + '*.es6', ['babel']);
});

gulp.task('default', ['babel', 'watch']);
