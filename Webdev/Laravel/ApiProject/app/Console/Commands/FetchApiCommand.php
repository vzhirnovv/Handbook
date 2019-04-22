<?php

namespace App\Console\Commands;

use Artisan;
use Illuminate\Console\Command;
use GuzzleHttp\Client;
use App\Post;

class FetchApiCommand extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'fetch:api';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Fetch API data from jsonplaceholder.';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    protected $client;

    public function __construct(Client $client)
    {
        parent::__construct();
        $this->client = $client;
    }

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        Artisan::call('migrate:refresh');

        $posts = $this->fetchPost();

        $bar = $this->output->createProgressBar(count($posts));

        foreach ($posts as $post) {
            Post::create([
          'title' => $post->title,
          'body' => $post->body,
        ]);
          $bar->advance();
        }
        $bar->finish();

        $this->info("\nApi seeding complete.");
    }

    protected function fetchPost()
    {
      $response = $this->client->request('GET', 'https://jsonplaceholder.typicode.com/posts');
      return json_decode($response->getBody());
    }
}
