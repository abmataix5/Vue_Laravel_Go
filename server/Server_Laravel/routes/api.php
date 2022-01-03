<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\AuthController;
use App\Http\Controllers\CourtController;
use App\Http\Controllers\UserController;
/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/


 Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
    return $request->user();
});

Route::resource('courts', CourtController::class);
Route::resource('users', UserController::class);

Route::post('/login', [AuthController::class, 'login']);

//NEW ADD FOR LOGIN
// Route::group(['prefix' => 'auth', 'middleware' => 'jwt.verify'], function () {
//     Route::post('logout', 'UserController@logout')->name('users.logout');
//     Route::put('update', 'UserController@update')->name('users.update');
//     Route::post('destroy', 'UserController@destroy')->name('users.destroy');
//     Route::post('getuser', 'UserController@getUser')->name('users.getuser');
// });

// Route::post('register', 'UserController@store')->name('users.store');
// Route::post('login', 'UserController@login')->name('users.login');

Route::group([
    'middleware' => 'api',
    'prefix' => 'auth'

], function ($router) {
    // Route::post('login', [AuthController::class, 'login']);
    Route::post('/register', [AuthController::class, 'register']);
    Route::post('/logout', [AuthController::class, 'logout']);
    Route::post('/refresh', [AuthController::class, 'refresh']);
    Route::get('/user-profile', [AuthController::class, 'userProfile']);   
     
});
