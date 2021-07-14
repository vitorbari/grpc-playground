<?php

use Calculator\FactorialRequest;

require "vendor/autoload.php";

$client = new Calculator\CalculatorServiceClient(
    "0.0.0.0:50051",
    ['credentials' => Grpc\ChannelCredentials::createInsecure()]
);

$numbers = range(1, 20);

foreach ($numbers as $number) {
    $request = new FactorialRequest();
    $request->setNumber($number);

    list($response, $status) = $client->Factorial($request, [], ['timeout' => 2000])->wait();

    if ($status->code !== Grpc\STATUS_OK) {
        echo sprintf(
            "Error when calculating the factorial of: %d. Error: %d, %s\n",
            $request->getNumber(),
            $status->code,
            $status->details
        );
        continue;
    }
    echo sprintf(
        "Factorial of %d = %d\n",
        $request->getNumber(),
        $response->getResult()
    );
}
