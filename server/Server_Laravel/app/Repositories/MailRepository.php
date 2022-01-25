<?php

namespace App\Repositories;

use App\Helpers\FileUploader;
use App\Models\User;
use App\Mail\TestEmail;
use App\Http\Controllers\PostController;

class MailRepository{


    public function sendMail(array $data)
    {
        ///DEBUG
        $out = new \Symfony\Component\Console\Output\ConsoleOutput();
        $out->writeln("---------------REPOSITORY SEND MAIL-------------------");
        $out->writeln($data['email']);
        $input = ['message' => 'Registrado en la plataforma PADEL-COURT'];
        // Mail::to($data['email'])->send(new TestEmail($input));

    }


}