var gulp = require('gulp');
var dotenv = require('dotenv');

gulp.task('environment', function() {
  process.env.NODE_ENV = process.env.ENVIRONMENT || 'development';
  dotenv.load();
});

require('./tasks/development');
