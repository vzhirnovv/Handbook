@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row">
        <div class="col-md-8 col-md-offset-2">
            <div class="panel panel-default">
                <div class="panel-heading">Posts</div>

                <div class="panel-body">
                  @foreach($posts as $post)
                    <h1>{{$post->title}}</h1>
                    <h2>{{$post->user_id}}</h2>
                    <h3>{{$post->user->email}}</h3>
                  @endforeach
                </div>
            </div>
        </div>
    </div>
</div>
@endsection
