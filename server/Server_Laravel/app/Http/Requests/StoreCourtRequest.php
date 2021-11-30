<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class StoreCourtRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     *
     * @return bool
     */
    public function authorize()
    {
        return false;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array
     */
    public function rules()
    {
        return [
            "name" => ["required"],
            "schedule" => ["required"],
            "date" => ["required"],
            "A_drive" => ["required"],
            "A_reves" => ["required"],
            "B_drive" => ["required"],
            "B_reves" => ["required"],
        ];
    }
}
