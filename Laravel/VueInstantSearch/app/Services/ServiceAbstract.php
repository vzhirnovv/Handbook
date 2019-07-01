<?php

namespace App\Services;

use GuzzleHttp\Client;

abstract class ServiceAbstract
{
  abstract public function get();

  public function __construct(Client $client)
  {
    $this->client = $client;
  }

}
