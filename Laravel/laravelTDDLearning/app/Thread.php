<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class Thread extends Model
{
    protected $fillable = [
      'user_id', 'title', 'body'
    ];

    public static function boot()
    {
      parent::boot();

      static::creating(function ($thread) {
        $thread->user_id = auth()->id();
      });
    }

    public function user()
    {
      return $this->belongsTo(User::class);
    }
}
