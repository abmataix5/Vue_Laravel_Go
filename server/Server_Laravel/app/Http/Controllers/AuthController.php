<?php

namespace App\Http\Controllers;
use Illuminate\Http\Request;

use Illuminate\Support\Facades\Auth;
use App\Models\User;
use Validator;

use App\Repositories\AuthRepository;
use App\Traits\ApiResponseTrait;

class AuthController extends Controller
{

    public $authRepository;
    use ApiResponseTrait;


    /**
     * Create a new AuthController instance.
     *
     * @return void
     */
    public function __construct(AuthRepository $authRepository) {
        $this->middleware('auth:api', ['except' => ['login', 'register','isAdmin']]);
        $this->authRepository = $authRepository;
    }

    /**
     * Get a JWT via given credentials.
     *
     * @return \Illuminate\Http\JsonResponse
     */
    public function login(Request $request){

        ///DEBUG
        $out = new \Symfony\Component\Console\Output\ConsoleOutput();
        $out->writeln("---------------LOGIN-------------------");
        $out->writeln("ENTRA LOGIN");
      
        $out->writeln("---------------LOGIN-------------------");
        /////DEBUG

    	$validator = Validator::make($request->all(), [
            'email' => 'required|email',
            'password' => 'required|string|min:6',
        ]);

        if ($validator->fails()) {
            return response()->json($validator->errors(), 422);
        }
     /*    $out->writeln($token); */
        if (! $token = auth()->attempt($validator->validated())) {
      
            return response()->json(['error' => 'Unauthorized'], 401);
        }

        return $this->createNewToken($token);
    }

    /**
     * Register a User.
     *
     * @return \Illuminate\Http\JsonResponse
     */
    public function register(Request $request) {
        // console.log("ENTRA REGISTER");
        $validator = Validator::make($request->all(), [
            'name' => 'required|string|between:2,100',
            'email' => 'required|string|email|max:100|unique:users',
            'password' => 'required|string|confirmed|min:6',
        ]);

        if($validator->fails()){
            return response()->json($validator->errors()->toJson(), 400);
        }

        $user = User::create(array_merge(
                    $validator->validated(),
                    ['password' => bcrypt($request->password)]
                ));

        return response()->json([
            'message' => 'User successfully registered',
            'user' => $user
        ], 201);
    }


    /**
     * Log the user out (Invalidate the token).
     *
     * @return \Illuminate\Http\JsonResponse
     */
    public function logout() {
        auth()->logout();

        return response()->json(['message' => 'User successfully signed out']);
    }

    /**
     * Refresh a token.
     *
     * @return \Illuminate\Http\JsonResponse
     */
    public function refresh() {
        return $this->createNewToken(auth()->refresh());
    }

    /**
     * Get the authenticated User.
     *
     * @return \Illuminate\Http\JsonResponse
     */
    public function userProfile() {
          ///DEBUG
          $out = new \Symfony\Component\Console\Output\ConsoleOutput();
          $out->writeln("---------------USERPROFILE-------------------");
          /////DEBUG
        return response()->json(auth()->user());
    }

    public function isAdmin(Request $request) {  //go consulta este endpoint para saber el tipos de usuario auth(). 
        ///DEBUG
        $out = new \Symfony\Component\Console\Output\ConsoleOutput();
        $out->writeln("---------------IS_ADMIN-------------------");
        /////DEBUG

        $data=$this->authRepository->isAdmin($request->email);
        if(is_null($data)){
            return self::apiResponseNotFound($request->email, 'Usuario no encontrado');
        }else{
            return $data;
        }    
    }



    /**
     * Get the token array structure.
     *
     * @param  string $token
     *
     * @return \Illuminate\Http\JsonResponse
     */
    protected function createNewToken($token){
         ///DEBUG
         $out = new \Symfony\Component\Console\Output\ConsoleOutput();
         $out->writeln("---------------LOGIN-------------------");
         $out->writeln($token);
       
         $out->writeln("---------------LOGIN-------------------");
         /////DEBUG
        return response()->json([
            'access_token' => $token,
            'token_type' => 'bearer',
            'expires_in' => auth()->factory()->getTTL() * 60,
            'user' => auth()->user()
        ]);
    }

}