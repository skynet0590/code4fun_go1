/*jshint esversion: 6 */

const autoprefixer = require('gulp-autoprefixer');
const gulp = require('gulp');
const sass = require('gulp-sass');
const sourcemaps = require('gulp-sourcemaps');

const style_main = './web/src/sass/main.scss';
const style_files = './web/src/sass/**/*.scss';

gulp.task('all', (done) => {
	gulp.src([style_main])
		.pipe(sourcemaps.init())
		.pipe(sass({
			outputStyle:'compressed',
			includePaths: [
				require('node-normalize-scss').includePaths,
				"node_modules"
			]
		})).on('error', sass.logError)
		.pipe(autoprefixer())
		.pipe(sourcemaps.write('./'))
		.pipe(gulp.dest('./web/public/css/'))
		.on('end', done);

	gulp.src([
		"./node_modules/jquery/dist/jquery.min.js",
		"./node_modules/bootstrap/dist/js/bootstrap.min.js"
	])
		.pipe(gulp.dest('./web/public/js/'))
		.on('end', done);
})

gulp.task('styles', (done)=>{
	gulp.src([style_main])
		.pipe(sourcemaps.init())	
		.pipe(sass({
			outputStyle:'compressed',
			includePaths: [
				require('node-normalize-scss').includePaths,
				"node_modules"
			]
		})).on('error', sass.logError)
		.pipe(autoprefixer())
		.pipe(sourcemaps.write('./'))
		.pipe(gulp.dest('./web/public/css/'))
		.on('end', done);
});

gulp.task('vendor-scripts', (done)=>{
	gulp.src([
		"./node_modules/jquery/dist/jquery.min.js",
		"./node_modules/bootstrap/dist/js/bootstrap.min.js"
		])
		.pipe(gulp.dest('./web/public/imported/js/'))
		.on('end', done);

});

gulp.task('watch', gulp.series(['all'], ()=>{
	gulp.watch('./src/sass/**/*.scss', gulp.series('styles'));
	console.log('watching files for changes');
}));
