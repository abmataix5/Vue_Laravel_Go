<?php

namespace App\Http\Requests;

use App\Models\User;
use Illuminate\Foundation\Http\FormRequest;
use Illuminate\Validation\ValidationException;
use Illuminate\Contracts\Validation\Validator;
use Illuminate\Http\Exceptions\HttpResponseException;

class StoreUserRequest extends FormRequest
{
    
    /**
     * Determine if the user is authorized to make this request.
     *
     * @return bool
     */
    public function authorize()
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array
     */
    public function rules()
    {
         ///DEBUG
         $out = new \Symfony\Component\Console\Output\ConsoleOutput();
         $out->writeln("---------------REQUEST-------------------");
         /////DEBUG
        return [
            "name" => ["required"],
            "email" => ["required"],
            "password" => ["required"],
        ];
    }

    public function withValidator($validator)
    {
        if ($validator->fails()) {
            ///DEBUG
         $out = new \Symfony\Component\Console\Output\ConsoleOutput();
         $out->writeln("---------------FAILS REQUEST-------------------");
         $out->writeln($this);
         /////DEBUG
            throw new HttpResponseException(response()->json([
                'msg'    => 'Rellene los campos correctamente',
                'status' => false,
                'errors' => $validator->errors(),
                'url'    => route('users.store')
            ], 400));
       }
    }
}