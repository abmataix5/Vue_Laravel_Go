<?php

namespace App\Http\Resources;

use Illuminate\Http\Resources\Json\JsonResource;

class UserResource extends JsonResource
{
    /**
     * Transform the resource into an array.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return array|\Illuminate\Contracts\Support\Arrayable|\JsonSerializable
     */
    public function toArray($request)
    {
        return [
            'id' => $this->id,
            'name' => $this->name,
            'lastname'=> $this->lastname,
            'email' => $this->email,
            'phone' => $this->phone,
            'position' => $this->position,
            'admin' => $this->admin,
            'image' => $this->image,
        ];
    }
}