<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Http\Requests\PostRequest;
use App\Post;
class PostController extends Controller
{
    public function index()
    {
      return view('posts/index');
    }

    public function create()
    {
      return view('/posts/create');
    }

    public function store(PostRequest $request)
    {
        Post::create([
          'title' => $request->title,
          'body' => $request->body
        ]);

        return redirect('/posts/create');
    }
}
