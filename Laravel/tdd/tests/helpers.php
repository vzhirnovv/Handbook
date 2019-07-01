<?php

function create($model)
{
    return factory($model)->create();
}

function make($model)
{
    return factory($model)->make();
}

function raw($model)
{
    return factory($model)->raw();
}
