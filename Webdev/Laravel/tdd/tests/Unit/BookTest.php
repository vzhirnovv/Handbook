<?php

namespace Tests\Unit;

use Tests\TestCase;
use Illuminate\Foundation\Testing\WithFaker;
use Illuminate\Foundation\Testing\RefreshDatabase;

class BookTest extends TestCase
{
    use RefreshDatabase; //VERY NECESSARY!!!!
    /**
     * A basic test example.
     *
     * @return void
     */
    public function test_a_book_belongs_to_many_authors()
    {
        $author = create('App\Author');
        // $author = factory('App\Author')->create();

        $book = factory('App\Book')->create();
        $book->authors()->syncWithoutDetaching([$author->id]);

        $this->assertDatabaseHas('author_book', [
            'author_id' => $author->id,
            'book_id' => $book->id,
        ]);
    }
}
