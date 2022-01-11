<?php

namespace App\Repositories;

use App\Models\User;

class AuthRepository{


    public function isAdmin($email)
    {

         ///DEBUG
         $out = new \Symfony\Component\Console\Output\ConsoleOutput();
         $out->writeln("---------------IS_ADMIN_REPOSITORY-------------------");
         $out->writeln($email);
         $user= User::where('email',$email)->get();
        //  $user= User::find($email);
         $out->writeln("--------------- valor email -------------------");
         $out->writeln($user);
         if($user){
             return $user;
         }else{
             return null;
         } 

    }


}