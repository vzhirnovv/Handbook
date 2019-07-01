<?php
  function create($class, $attributes = [], $times = null)
  {
    return factory($class, $times)->create($attributes);
  }

  function make($class, $attributes = [], $times = null)
  {
    return factory($class)->make($attributes);
  }

// GLOBAL HELPER PHP_Token_FUNCTION
// check composer.json -> autoload-dev -> files
?>
