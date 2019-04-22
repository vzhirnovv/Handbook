@extends('layouts.master')
@section('content')
      <div class="col-sm-8 blog-main">
        <h1>Create a post</h1>
      <form action="/posts/store" method="post">
          {{csrf_field()}}
          <div class="form-group">
              <input type="text" name="title" class="form-control">
          </div>

          <div class="form-group">
              <textarea name="body" class="form-control"></textarea>
          </div>

          <button type="submit" class="btn btn-primary" name="button">Publish</button>

      </form>

      @if ($errors->any())
          <div class="alert alert-danger">
            <ul>
                @foreach ($errors->all() as $error)
                  <li>{{ $error }}</li>
                @endforeach
            </ul>
          </div>
      @endif
    </div>
@endsection
