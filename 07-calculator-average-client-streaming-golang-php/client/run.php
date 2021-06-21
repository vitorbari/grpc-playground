<?php

require "vendor/autoload.php";

use Calculator\AverageRequest;

$client = new Calculator\CalculatorServiceClient(
    "0.0.0.0:50051",
    ['credentials' => Grpc\ChannelCredentials::createInsecure()]
);

$call = $client->Average();

$numbers = [1,2,3,4];

echo "Numbers: ";
print_r($numbers);

foreach ($numbers as $number) {
    $request = (new AverageRequest())
        ->setNumber($number);

    $call->write($request);
}

list($response, $status) = $call->wait();

echo sprintf("Response from RPC - Average = %.1f\n", $response->getResult());
