<?php

namespace App\Http\Controllers\Api;


use Illuminate\Http\Request;
use App\Http\Controllers\Controller;
use GuzzleHttp\Client;
use App\Post;

class HomeController extends Controller
{
    protected $client;

    public function __construct(Client $client)
    {
        $this->client = $client;
    }


    public function index()
    {
        // $response = Post::get();
        $response = $this->client->request('GET', 'https://jsonplaceholder.typicode.com/posts');
        $response = json_decode($response->getBody());
        return view('posts.index', compact('response'));
    }

    public function show($id)
    {
      $response = $this->client->request('GET', 'https://jsonplaceholder.typicode.com/posts/' . $id);

      $response = json_decode($response->getBody());

      return view('posts.show', compact('response'));
    }
}
