@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row justify-content-center">
        <div class="col-md-8">
            <div class="card" style="margin-top:30px">

                <div class="card-header">
                  <a href="/api/responses/{{$response->id}}">
                    {{$response->title}}
                  </a>
                </div>

                <div class="card-body">
                  {{$response->body}}
                </div>
            </div>
        </div>
    </div>
</div>
@endsection
