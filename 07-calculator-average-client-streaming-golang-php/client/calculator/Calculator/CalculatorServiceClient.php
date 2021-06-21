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
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\ClientStreamingCall
     */
    public function Average($metadata = [], $options = []) {
        return $this->_clientStreamRequest('/calculator.CalculatorService/Average',
        ['\Calculator\AverageResponse','decode'],
        $metadata, $options);
    }

}
