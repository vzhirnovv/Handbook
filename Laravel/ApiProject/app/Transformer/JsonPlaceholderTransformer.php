<?php

namespace App\Transformer;

use App\Transformer\TransformerAbstract;

class JsonPlaceholderTransformer extends TransformerAbstract
{
  public function transform($payload)
  {
    return [
      'phone' => $payload->phone,
      'website' => $payload->website,
    ];
  }
}
