<?php

namespace App\Http\Controllers;


use Illuminate\Http\Request;
use App\Phone;

class PhoneController extends Controller
{
    public function index()
    {
      $phones = Phone::all();
      return view('phone')->with('phones', $phones);
    }

    public function create()
    {
      return view('form');
    }

    public function store(Request $request,Phone $phone)
    {
      $this->validate($request,[
        'phone_number' => 'required'
      ],[
        'phone_number.required' => 'Enter a correct phone number'
      ]);

      $phone->user()->create([
        'phone_number' => $request->input('phone_number'),
      ]);
    }


}
