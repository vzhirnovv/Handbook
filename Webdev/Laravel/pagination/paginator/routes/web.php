<?php

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/
use App\post;

Route::get('/', function () {
    $posts = post::paginate(10);
    
    return view('post',compact('posts'));

});

Auth::routes();

Route::get('/home', 'HomeController@index')->name('home');
