<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: calculator.proto

namespace Calculator;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>calculator.PrimeNumberDecompositionRequest</code>
 */
class PrimeNumberDecompositionRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 number = 1;</code>
     */
    protected $number = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $number
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Calculator::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 number = 1;</code>
     * @return int
     */
    public function getNumber()
    {
        return $this->number;
    }

    /**
     * Generated from protobuf field <code>int32 number = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setNumber($var)
    {
        GPBUtil::checkInt32($var);
        $this->number = $var;

        return $this;
    }

}

