<?php

namespace App\Http\Controllers\Api;

use Illuminate\Http\Request;
use App\Http\Controllers\Controller;
use GuzzleHttp\Client;

class TransformerController extends Controller
{
    public function __construct(Client $client)
    {
        $this->client = $client;
    }

    public function index()
    {
        // $response = Post::get();
        $response = $this->client->request('GET', 'https://jsonplaceholder.typicode.com/users');
        dd(json_decode($response->getBody()));

        // return array_map(function ($user) {
        //     return [
        //     'phone' => $user->phone,
        //     'website' => $user->website,
        //   ];
        // }, $response);

        return (new JsonPlaceholderTransformer($users))->create();
    }
}
