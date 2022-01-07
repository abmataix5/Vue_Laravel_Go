<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

use App\Http\Resources\UserResource;
use App\Http\Resources\UserCollection;
use App\Http\Requests\StoreUserRequest; 
use App\Models\User;

class UserController extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index(Request $request)
    {
        ///DEBUG
        $out = new \Symfony\Component\Console\Output\ConsoleOutput();
        $out->writeln("---------------USER-------------------");
        /////DEBUG
        // return UserResource::collection(User::get());
        return new UserCollection(User::get());
        //
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(StoreUserRequest $request)
    {
         ///DEBUG
         $out = new \Symfony\Component\Console\Output\ConsoleOutput();
         $out->writeln("---------------USER_STORE-------------------");
         $out->writeln($request);
         /////DEBUG
         $user = new User;
        //  $user->id = $request->id;
         $user->name = $request->name;
         $user->email = $request->email;
         $user->password = $request->password;
         $user->save();
         return response()->json([
             "message" => "Socio registrado correctamente"
         ], 201);

    }

    /**
     * Display the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        ///DEBUG
        $out = new \Symfony\Component\Console\Output\ConsoleOutput();
        $out->writeln("---------------USER-------------------");
        /////DEBUG
        return UserResource::make(User::where('id', $id)->firstOrFail());
        //
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function edit($id)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, $id)
    {

        if (user::where('id', $id)->exists()) {
            $user = User::find($id);
            $user->id = $request->id;
            $user->name = $request->name;
            $user->email = $request->email;
          
            $user->save();
            return response()->json([
              "message" => "Socio actualizado correctamente"
            ], 200);
          } else {
            return response()->json([
              "message" => "court not found"
            ], 404);
          }

        //
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
        //
        ///DEBUG
        $out = new \Symfony\Component\Console\Output\ConsoleOutput();
        $out->writeln("---------------DELETE USER-------------------");
        /////DEBUG
        
        if(User::where('id', $id)->exists()) {
            $user = User::find($id);
            $user->delete();
            return response()->json([
              "message" => "usuario eliminado"
            ], 202);
          } else {
            return response()->json([
              "message" => "usuario no encontrado"
            ], 404);
          }
    }
}
