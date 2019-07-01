<?php

use Illuminate\Http\Request;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::middleware('auth:api')->get('/user', function (Request $request) {
    return $request->user();
});

// Show all lessons
Route::get('lessons', 'LessonController@index');

// Show single lesson
Route::get('lessons/{id}', 'LessonController@show');

// Post a lesson
Route::post('lesson', 'LessonController@store');

// Update a lesson
Route::put('lesson', 'LessonController@store');

// Delete a lesson
Route::delete('lessons/{id}', 'LessonController@destroy');







