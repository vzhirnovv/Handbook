<?php

use Faker\Generator as Faker;

$factory->define(App\Lesson::class, function (Faker $faker) {
    return [
        'url' => $faker->url(),
        'title' => $faker->text(5),
        'description' => $faker->text(10),
    ];
});
