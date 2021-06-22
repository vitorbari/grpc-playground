<?php

require "vendor/autoload.php";

use Calculator\MaximumRequest;

$client = new Calculator\CalculatorServiceClient(
    "0.0.0.0:50051",
    ['credentials' => Grpc\ChannelCredentials::createInsecure()]
);

$call = $client->Maximum();

$numbers = [1,5,3,6,2,20];

echo "Request numbers: ";
print_r($numbers);

foreach ($numbers as $number) {
    $request = (new MaximumRequest())
        ->setNumber($number);

    $call->write($request);
}
$call->writesDone();

/** @var Calculator\MaximumResponse $response */
while ($response = $call->read()) {
    echo sprintf("Current Maximum: %d\n", $response->getMaximum());
}
