<?php

namespace Tests\Feature;

use Tests\TestCase;
use Illuminate\Foundation\Testing\WithFaker;
use Illuminate\Foundation\Testing\RefreshDatabase;

class CreateThreadTest extends TestCase
{

    use RefreshDatabase;

    //Validation
    /** @test */
    public function a_thread_requires_a_title()
    {

      $this->withExceptionHandling();

      $this->signIn();

      $thread = create('App\Thread');

      dd($thread);

      $this->post('/threads/store')
        ->assertSessionHasErrors('title');
    }
}
