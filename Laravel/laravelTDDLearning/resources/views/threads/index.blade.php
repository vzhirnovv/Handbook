@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row">
        <div class="col-md-8 col-md-offset-2">
            <div class="panel panel-default">
                <div class="panel-heading">Dashboard</div>

                <div class="panel-body">
                  @foreach ($threads as $thread)
                    <h4>{{$thread->user_id}}</h4>
                    <h4>{{$thread->title}}</h4>
                    <h6>{{$thread->body}}</h6>
                  @endforeach
                </div>
            </div>

        <div class="row">
          <form action="/threads/store" method="post">
            {{csrf_field()}}

            <div class="form-group">
              <input type="text" name="title">
            </div>

            <div>
              <input type="text" name="body">
            </div>

            <button type="submit" class='btn btn-primary'>send shit</button>
          </form>
        </div>
      </div>
    </div>
</div>
@endsection
