<?php

namespace Tests\Unit;

use Tests\TestCase;
use Illuminate\Foundation\Testing\WithFaker;
use Illuminate\Foundation\Testing\RefreshDatabase;

class ThreadTest extends TestCase
{
    // A thread belongs to a user

    /** @test */
    public function a_thread_belongs_to_a_user()
    {
      $thread = factory('App\Thread')->create();


      $this->assertInstanceOf('App\User', $thread->user);

      // dd(get_class($thread->user));
    }
}
