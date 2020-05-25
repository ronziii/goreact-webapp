var gulp = require('gulp');
var runSequence = require('run-sequence');
var shell = require('gulp-shell');
var server = require('./server');

var developmentServer;

gulp.task('app:install', ['environment'], function() {
  return gulp.src('').pipe(shell('go install', { cwd: './cmd/ictlife_infra_interview_may_2020' }));
});

gulp.task('server', function() {
  runSequence('environment', 'app:install', function() {
    developmentServer = server.start({silent: false});

    gulp.watch(['*.go', 'cmd/**/*.go', 'app/**/*.go', 'app/**/**/*.go'], ['server:restart']);
  });
});

gulp.task('server:restart', ['environment', 'app:install'], function(done) {
  if (!developmentServer) {
    return;
  }

  developmentServer.on('close', function() {
    developmentServer = server.start({silent: false});
    done();
  });
  developmentServer.kill();
});

gulp.task('dropdb', ['environment'], () => {
  return gulp.src('')
    .pipe(shell('docker-compose exec -T postgres psql -c "DROP DATABASE IF EXISTS ictlife_infra_interview_may_2020;"'));
});

gulp.task('createdb', ['environment'], () => {
  return gulp.src('')
    .pipe(shell('docker-compose exec -T postgres psql -c "CREATE DATABASE ictlife_infra_interview_may_2020;"'));
});

gulp.task('migratedb', ['environment'], () => {
  return gulp.src('')
    .pipe(shell('goose -dir \'app/db/migrations\' postgres $DATABASE_URL up'));
});
