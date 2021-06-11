<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Calculator;

/**
 */
class CalculatorServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Calculator\PrimeNumberDecompositionRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\ServerStreamingCall
     */
    public function GetPrimeNumberDecomposition(\Calculator\PrimeNumberDecompositionRequest $argument,
      $metadata = [], $options = []) {
        return $this->_serverStreamRequest('/calculator.CalculatorService/GetPrimeNumberDecomposition',
        $argument,
        ['\Calculator\PrimeNumberDecompositionResponse', 'decode'],
        $metadata, $options);
    }

}
