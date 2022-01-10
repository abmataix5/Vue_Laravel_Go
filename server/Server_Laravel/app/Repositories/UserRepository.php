<?php

namespace App\Repositories;

use App\Models\User;

class UserRepository{


    public function all()
    {
        return User::get();
    }

    public function create(array $data)
    {
        return User::create($data);
    }

    public function update($id, array $data)
    {
        $user= User::find($id);
        if($user){
            return $user->update($data);
        }else{
            return null;
        }     
    }

    public function delete($id){

        $user= User::find($id);
        if($user){
            return $user->delete($user);
        }else{
            return null;
        }
    }

}