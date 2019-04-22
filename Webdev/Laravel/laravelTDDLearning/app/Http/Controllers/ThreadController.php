<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Thread;

class ThreadController extends Controller
{
  public function index() {
    $threads = Thread::all();

    return view('threads.index', compact('threads'));
  }

  public function store(Request $request) {
    Thread::create([
      'title' => $request->title,
      'body' => $request->body,
      // 'user_id' => auth()->id()
    ]);

    return back();
  }
}
