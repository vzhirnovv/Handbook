@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row">
      <div class="col-md-4">
        <li><a href="/listings/?sort=desc">Desc</a></li>
        <li><a href="/listings/?sort=asc">Asc</a></li>
      </div>
        <div class="col-md-8 col-md-offset-2">
            <div class="panel panel-default">
                <div class="panel-heading">Lisitngs</div>
                <div class="panel-body">
                    @foreach ($listings as $listing)
                      <h1>{{$listing->title}}</h1>
                      <h2>{{$listing->type}}</h2>
                      <h3>{{$listing->created_at->diffForHumans()}}</h3>
                      <h3>{{$listing->created_at->format('Y-m-d')}}</h3>
                      <hr>
                    @endforeach
                </div>
            </div>
        </div>
    </div>
</div>
@endsection
