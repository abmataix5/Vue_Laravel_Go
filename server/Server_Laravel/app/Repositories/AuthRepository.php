<?php

namespace App\Repositories;

use App\Models\User;
use App\Models\Worker;

class AuthRepository{


    public function isAdmin($email) //buscar user auth desde Go
    {
         $user= User::where('email',$email)->get(); // buscamos usuario por email.
       
         if($user){
             return $user;
         }else{
             return null;
         } 

    }


}