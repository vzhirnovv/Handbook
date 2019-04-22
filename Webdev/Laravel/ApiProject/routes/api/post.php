<?php

Route::get('/posts', 'Api\HomeController@index');
Route::get('/posts/{id}', 'Api\HomeController@show');
