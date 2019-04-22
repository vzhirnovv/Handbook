<?php

use Faker\Generator as Faker;

/* @var Illuminate\Database\Eloquent\Factory $factory */

$factory->define(App\Listing::class, function (Faker $faker) {
    return [
        'title' => $faker->sentence(4),
        'type' => 'male',
    ];
});
