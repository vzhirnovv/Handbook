@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row">
        <div class="col-md-8 col-md-offset-2">
            <div class="panel panel-default">
                <div class="panel-heading">Phone Number</div>

                  <div class="panel-body">
                    @foreach ($phones as $phone)
                      {{$phone->phone_number}} of {{$phone->user->name}}
                      <br>
                      <hr>
                    @endforeach
                  </div>

                </div>
            </div>
        </div>
    </div>
</div>
@endsection
