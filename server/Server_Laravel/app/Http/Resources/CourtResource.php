<?php

namespace App\Http\Resources;

use Illuminate\Http\Resources\Json\JsonResource;

class CourtResource extends JsonResource
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
            'schedule' => $this->schedule,
            'date' => $this->date,
            'A_drive' => $this->A_drive,
            'A_reves' => $this->A_reves,
            'B_drive' => $this->B_drive,
            'B_reves' => $this->B_reves
        ];
    }
}
