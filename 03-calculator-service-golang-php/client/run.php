<?php

use Calculator\Addition;
use Calculator\AdditionRequest;

require "vendor/autoload.php";

$client = new Calculator\CalculatorServiceClient(
    "0.0.0.0:50051",
    ['credentials' => Grpc\ChannelCredentials::createInsecure()]
);

$addition = (new Addition())
    ->setAugend(7)
    ->setAddend(5);

$request = new AdditionRequest();
$request->setAddition($addition);

list($response, $status) = $client->Add($request)->wait();

if ($status->code !== Grpc\STATUS_OK) {
    echo sprintf("ERROR: %d, %s", $status->code, $status->details).PHP_EOL;
    exit(1);
}
echo sprintf("Response from Calculator RPC: %d", $response->getResult()).PHP_EOL;
