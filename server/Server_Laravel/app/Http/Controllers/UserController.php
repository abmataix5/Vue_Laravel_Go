<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

use App\Http\Resources\UserResource;
use App\Http\Resources\UserCollection;
use App\Http\Requests\StoreUserRequest; 

use App\Traits\ApiResponseTrait;
use App\Repositories\UserRepository;
// use App\Repositories\MailRepository;

use App\Models\User;


class UserController extends Controller
{

    use ApiResponseTrait;
    public $userRepository;
    // public $mailRepository;
    public function __construct(UserRepository $userRepository)
    // public function __construct(UserRepository $userRepository,MailRepository $mailRepository)
    {
        $this->userRepository = $userRepository;
        // $this->mailRepository = $mailRepository;
     
    }

    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index(Request $request)
    {
        $data=$this->userRepository->all();
    
        return self::apiResponseSuccess(new UserCollection($data), 'Listado Socios');
    
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
        try{
         
            // $data= $this->mailRepository->sendMail($request->all());
            $data = $this->userRepository->create($request->all());

            return self::apiResponseSuccess($data, 'Socio registrado correctamente');

        }catch(\Throwable|\Exception $e){
            return self::apiResponseServerError($e, 'Error server interno.');
        }
    }

    /**
     * Display the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        return UserResource::make(User::where('id', $id)->firstOrFail());
    
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
    public function update( Request $request, $id)
    {
        try{
        ///DEBUG
        $out = new \Symfony\Component\Console\Output\ConsoleOutput();
        $out->writeln("---------------USER_UPDATE-------------------");
        // $out->writeln($request);
        // $out->writeln($id);
        $file = $request->query('name');
        // $file = $request->input();
        $out->writeln($file);


        $data = $this->userRepository->update($id,$request->all());

        if(is_null($data)){
            return self::apiResponseNotFound($id, 'usuario no encontrado');
        } else {
            return self::apiResponseSuccess($id, 'Socio actualizado correctamente');
        }
    }catch(\Throwable|\Exception $e){
        return self::apiResponseServerError($e, 'Error server interno.');
    }
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
        $data = $this->userRepository->delete($id);
        if(is_null($data)){
            return self::apiResponseNotFound($id, 'Socio no encontrado');
        } else {
            return self::apiResponseSuccess($id, 'Socio eliminado correctamente');
        }
    }
}
