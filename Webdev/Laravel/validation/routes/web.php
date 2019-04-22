<?php

Route::get('/', 'PostController@index');
Route::get('/posts/create', 'PostController@create');
Route::post('/posts/store', 'PostController@store');

Auth::routes();

Route::get('/home', 'HomeController@index')->name('home');
