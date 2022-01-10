<?php

namespace App\Repositories;

use App\Models\Court;

class CourtRepository{


    public function all()
    {
        return Court::get();
    }

    public function create(array $data)
    {
        return Court::create($data);
    }

    public function find($id){
        return Court::find($id);
    }

    public function update($id, array $data)
    {
        $court= Court::find($id);
        if($court){
            return $court->update($data);
        }else{
            return null;
        }     
    }

    public function delete($id){

        $court= Court::find($id);
        if($court){
            return $court->delete($court);
        }else{
            return null;
        }
    }
}