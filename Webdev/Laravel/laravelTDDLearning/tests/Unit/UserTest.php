<?php

namespace Tests\Unit;

use Tests\TestCase;
use Illuminate\Foundation\Testing\WithFaker;
use Illuminate\Foundation\Testing\RefreshDatabase;

class UserTest extends TestCase
{
    //user has many threads

    /** @test */

    public function a_user_has_many_threads()
    {
      $user = create('App\User');

      // $thread = create('App\Thread');

      dd($user->name);

      $this->assertTrue($user->threads->contains($thread));
    }
}
