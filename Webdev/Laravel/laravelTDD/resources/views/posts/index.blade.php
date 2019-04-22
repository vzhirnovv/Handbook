@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row">
        <div class="col-md-8 col-md-offset-2">
            <div class="panel panel-default">
                <div class="panel-heading">index</div>

                <div class="panel-body">
                    @foreach ($posts as $post)
                      <a href="/posts/{{$post->id}}">
                        {{$post->title}}
                      </a>
                      {{$post->body}}
                    @endforeach
                </div>
            </div>
        </div>
    </div>
</div>
@endsection
