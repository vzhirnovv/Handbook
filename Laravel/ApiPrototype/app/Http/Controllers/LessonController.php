<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

use App\Lesson;
use App\Http\Requests;
use App\Http\Resources\Lesson as LessonResource;

class LessonController extends Controller
{
    public function index()
    {
        $lessons = Lesson::paginate(15);

        return LessonResource::collection($lessons);
    }

    public function store(Request $request)
    {
        $lesson = $request->isMethod('put') ? Lesson::findOrFail
        ($request->lesson_id) : new Lesson;
        $lesson->id = $request->input('lesson_id');
        $lesson->url = $request->input('url');
        $lesson->title = $request->input('title');
        $lesson->description = $request->input('description');

        if ($lesson->save()){
            return new LessonResource($lesson);
        }

    }

    public function show($id)
    {
        $lesson = Lesson::findOrFail($id);

        return new LessonResource($lesson);
    }

    public function destroy($id)
    {
        $lesson = Lesson::findOrFail($id);

        if ($lesson->delete()){
            return new LessonResource($lesson);
        }
    }
}
