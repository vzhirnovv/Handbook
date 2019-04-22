<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class Reply extends Model
{

    protected $fillable = [
      'user_id', 'body', 'thread_id'
    ];
}
