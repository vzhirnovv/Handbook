@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row justify-content-center">
        <div class="col-md-8">
          @foreach ($response as $post)
            <div class="card" style="margin-top:30px">

                <div class="card-header">
                  <a href="/api/posts/{{$post->id}}">
                    {{$post->title}}
                  </a>
                </div>

                <div class="card-body">
                  {{$post->body}}
                </div>
            </div>
          @endforeach
        </div>
    </div>
</div>
@endsection
