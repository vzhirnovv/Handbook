<?php

namespace Tests\Feature;

use Tests\TestCase;
use Illuminate\Foundation\Testing\WithFaker;
use Illuminate\Foundation\Testing\RefreshDatabase;


class ReadPostsTest extends TestCase
{
    use RefreshDatabase;

    protected $post;

    public function setUp()
    {
      parent::setUp();

      $this->post = factory('App\Post')->create();
    }

    /** @test */
    public function an_authenticated_user_can_read_all_posts()
    {

        $this->signIn();

        $this->get('/posts')
            ->assertSee($this->post->title)
            ->assertSee($this->post->body);
    }

    /** @test */
    public function an_authenticated_user_can_read_a_single_post()
    {
      $this->get('/posts/' . $this->post->id)
            ->assertSee($this->post->title)
            ->assertSee($this->post->body);
    }

    /** @test */
    public function an_authenticated_user_can_create_a_post()
    {
      $post = factory('App\Post')->create();
      dd($post);

      $response = $this->post('/posts', $post->toArray());

      $this->get($response->headers->get('Location'))
          ->assertSee($post->title);

      //dd($response->headers->get('Location'));
    }
}
