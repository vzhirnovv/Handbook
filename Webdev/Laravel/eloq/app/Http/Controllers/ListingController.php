<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Listing;

class ListingController extends Controller
{
    public function index(Request $request)
    {
      $listings = new Listing;
      //if ($request->has('type')){
      //  $listings = $listings->where('type', request('type'));
      //}

      //$listings = $listings->get();

      if ($request->has('sort')){
        $listings = $listings->orderBy('created_at', request('sort'));
      }

      $listings = $listings->get();

      return view('listings.index')->with('listings', $listings);
    }
}
