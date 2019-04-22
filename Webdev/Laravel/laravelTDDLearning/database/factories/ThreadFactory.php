<?php

use Faker\Generator as Faker;

$factory->define(App\Thread::class, function (Faker $faker) {
    return [
        'user_id' => function () {
          return factory('App\User')->create()->id;
        },
        'title' => $faker->word(),
        'body' => $faker->sentence(8)
    ];
});
