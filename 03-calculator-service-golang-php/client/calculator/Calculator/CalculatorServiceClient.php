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
     * @param \Calculator\AdditionRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function Add(\Calculator\AdditionRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/calculator.CalculatorService/Add',
        $argument,
        ['\Calculator\AdditionResponse', 'decode'],
        $metadata, $options);
    }

}
