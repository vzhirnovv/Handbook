<?php

namespace App\Http\Controllers;

use GuzzleHttp\Client;
use Illuminate\Http\Request;
use App\Services\LibrariesService;
use App\Services\Cache\RedisAdapter;

class LibraryController extends Controller
{
    protected $client;
    protected $cache;

    public function __construct(Client $client, RedisAdapter $cache)
    {
      $this->client = $client;
      $this->cache = $cache;
    }


    public function index()
    {
      $data = $this->cache->remember('libraries', 10, function(){
        return json_encode((new LibrariesService($this->client))->get());
      });

      return json_decode($data);
    }


}
