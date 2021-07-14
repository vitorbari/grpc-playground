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
     * @param \Calculator\FactorialRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function Factorial(\Calculator\FactorialRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/calculator.CalculatorService/Factorial',
        $argument,
        ['\Calculator\FactorialResponse', 'decode'],
        $metadata, $options);
    }

}
