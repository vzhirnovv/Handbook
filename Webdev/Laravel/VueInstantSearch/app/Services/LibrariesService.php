<?php

namespace App\Services;

use App\Services\ServiceAbstract;

class LibrariesService extends ServiceAbstract
{
  public function get()
  {
    $response = $this->client->request('GET', 'https://api.cdnjs.com/libraries');
    $response = json_decode($response->getBody());

    return $response->results;
  }
}
