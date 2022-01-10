<?php
 
namespace App\Traits;
 
use Illuminate\Http\Request;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Response;
 
trait ApiResponseTrait {
    
	public static function apiResponseSuccess($data, $message = "Success", $code = JsonResponse::HTTP_OK)
	{
         ///DEBUG
         $out = new \Symfony\Component\Console\Output\ConsoleOutput();
         $out->writeln("---------------TRAITS-------------------");
         /////DEBUG
         
		return response()->json([
            'status' => true,
            'message' => $message,
            'data' => $data,
        ], $code);
	}

    public static function apiResponseNotFound($data, $message = "Not found", $code = JsonResponse::HTTP_OK)
	{
            ///DEBUG
            $out = new \Symfony\Component\Console\Output\ConsoleOutput();
            $out->writeln("---------------TRAITS-------------------");
            /////DEBUG
		return response()->json([
            'status' => true,
            'message' => $message,
            'data' => $data,
        ], $code);
	}

    public static function apiResponseServerError($data, $message = "Internal Server Error")
	{
		return response()->json([
            'status' => false,
            'message' => $message,
            'data' => $data,
        ], Response::HTTP_INTERNAL_SERVER_ERROR);
	}

}