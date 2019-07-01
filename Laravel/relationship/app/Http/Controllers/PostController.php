<?php

namespace App\Http\Controllers;

use App\User;
use App\Post;
use Illuminate\Http\Request;

class PostController extends Controller
{
    public function index(Request $request)
    {
      $posts = Post::all();
      return view('users.posts.index')->with('posts', $posts);
    }
}
