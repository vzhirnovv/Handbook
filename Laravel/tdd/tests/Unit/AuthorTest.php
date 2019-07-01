<?php

namespace Tests\Unit;

use Tests\TestCase;
use Illuminate\Foundation\Testing\WithFaker;
use Illuminate\Foundation\Testing\RefreshDatabase;

class AuthorTest extends TestCase
{
    use RefreshDatabase;

    /**
     * A basic test example.
     *
     * @return void
     */
    public function test_an_author_belongs_to_many_books()
    {
        $book = factory('App\Book')->create();

        $author = factory('App\Author')->create();
        $author->books()->syncWithoutDetaching([$book->id]);

        $this->assertDatabaseHas('author_book', [
            'author_id' => $author->id,
            'book_id' => $book->id,
        ]);
    }
}
