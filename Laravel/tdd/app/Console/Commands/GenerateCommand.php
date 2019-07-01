<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;

class GenerateCommand extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'books:create';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Make some books';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        $books = factory('App\Book', 1000)->create();

        $bar = $this->output->createProgressBar(count($books));

        foreach ($books as $book) {
            $this->performTask($book);

            $bar->advance();
        }

        $bar->finish();
    }
}
