@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row">
        <div class="col-md-8 col-md-offset-2">
            <div class="panel panel-default">
                <div class="panel-heading">Create a phone number</div>

                  <div class="panel-body">
                    <form method="post" action="/create/phone">
                        {{csrf_field()}}
                        <div class="form-group">
                          <label for="phone_number">Phone Number</label>
                          <input type="text" class="form-control" name="phone_number">

                        </div>

                          <button type="submit" class="btn btn-primary">Submit</button>
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

                </div>
            </div>
        </div>
    </div>
</div>
@endsection
