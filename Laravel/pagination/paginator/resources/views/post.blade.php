@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row">
        <div class="col-md-8 col-md-offset-2">
            <div class="panel panel-default">
                <div class="panel-heading">Post</div>

                <div class="panel-body">
                  @foreach ($posts as $post)
                    {{$post->title}}
                    <br>
                    {{$post->body}}
                    <hr>
                  @endforeach

                  {{$posts->links()}}

                </div>
            </div>
        </div>
    </div>
</div>
@endsection
