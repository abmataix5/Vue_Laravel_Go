<?php

namespace App\Repositories;

use App\Models\User;

class AuthRepository{


    public function isAdmin($email)
    {
         $user= User::where('email',$email)->get(); // buscamos usuario por email.
       
         if($user){
             return $user;
         }else{
             return null;
         } 

    }


}