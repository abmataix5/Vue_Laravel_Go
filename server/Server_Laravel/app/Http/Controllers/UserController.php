<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

use App\Http\Resources\UserResource;
use App\Http\Resources\UserCollection;
use App\Http\Requests\StoreUserRequest; 

use App\Traits\ApiResponseTrait;

use App\Models\User;

class UserController extends Controller
{

    use ApiResponseTrait;


    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index(Request $request)
    {
        ///DEBUG
        $out = new \Symfony\Component\Console\Output\ConsoleOutput();
        $out->writeln("---------------USER LIST-------------------");
        /////DEBUG

        return self::apiResponseSuccess(new UserCollection(User::get()), 'Listado Socios');
    
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
         return self::apiResponseSuccess($data, 'Socio registrado correctamente');
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
            return self::apiResponseSuccess($id, 'Socio actualizado correctamente');
        } else {
            return self::apiResponseNotFound($id, 'usuario no encontrado');
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
        ///DEBUG
        $out = new \Symfony\Component\Console\Output\ConsoleOutput();
        $out->writeln("---------------DELETE USER-------------------");
        /////DEBUG
        
        if(User::where('id', $id)->exists()) {
            $user = User::find($id);
            $user->delete();

          return self::apiResponseSuccess($id, 'Socio eliminado correctamente');
          } else {

            return self::apiResponseNotFound($id, 'usuario no encontrado');
          }
    }
}
