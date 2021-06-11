<?php

require "vendor/autoload.php";

use Calculator\PrimeNumberDecompositionRequest;

$client = new Calculator\CalculatorServiceClient(
    "0.0.0.0:50051",
    ['credentials' => Grpc\ChannelCredentials::createInsecure()]
);

$request = (new PrimeNumberDecompositionRequest())
    ->setNumber(1200);

$call = $client->GetPrimeNumberDecomposition($request);

/** @var $factors Calculator\PrimeNumberDecompositionResponse[] */
$factors = $call->responses();

// an iterator over the server streaming responses
foreach ($factors as $key => $factor) {
    if ($key > 0) {
        echo ',';
    }
    echo $factor->getFactor();
}

echo PHP_EOL;
