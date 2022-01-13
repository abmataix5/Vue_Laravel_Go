<?php

namespace App\Repositories;

use App\Helpers\FileUploader;
use App\Models\User;
use App\Http\Controllers\PostController;

class UserRepository{


    public function all()
    {
        return User::get();
    }

    public function create(array $data)
    {
        ///DEBUG
        $out = new \Symfony\Component\Console\Output\ConsoleOutput();
        $out->writeln("---------------USER_CREATE-------------------");
     
        // upload image file
        if(isset($data['image'])){ 
            $out->writeln("---------------ENTRA-------------------");
            if($data['image'] != null && $data['image'] != '' && !is_string($data['image'])){
                $data['image'] = FileUploader::store($data['image'], $data['name'] ,'gallery/users'); 
            } 
        }else{
            $data['image']=NULL;
        }
        $out->writeln($data);
        $user = User::create($data);
        return $user;
    }

    public function update($id, array $data)
    {
        ///DEBUG
        $out = new \Symfony\Component\Console\Output\ConsoleOutput();
        $out->writeln("---------------USER_UPDATE-------------------");
        
        $out->writeln($data);

        $user= User::find($id);
        $out->writeln("---------------VALOR BUSQUEDA USER-------------------");
        $out->writeln($user);

        if($user){
            if(isset($data['image'])){ 
                if($data['image'] != null && $data['image'] != '' && !is_string($data['image'])){
                    $data['image']   = FileUploader::update('image', $data['image'], $data['title'] ,'gallery/users', $user->image);            
                } 
            }

            return $user->update($data);
        }else{
            return null;
        }     
    }

    public function delete($id){

        $user= User::find($id);
        if($user){
            FileUploader::delete('gallery/users/'.$user->image);
            return $user->delete($user);
        }else{
            return null;
        }
    }

}