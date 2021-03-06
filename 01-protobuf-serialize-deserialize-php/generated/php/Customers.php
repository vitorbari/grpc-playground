<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: customers.proto

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>Customers</code>
 */
class Customers extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .Customer customers = 1;</code>
     */
    private $customers;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Customer[]|\Google\Protobuf\Internal\RepeatedField $customers
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Customers::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .Customer customers = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getCustomers()
    {
        return $this->customers;
    }

    /**
     * Generated from protobuf field <code>repeated .Customer customers = 1;</code>
     * @param \Customer[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setCustomers($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Customer::class);
        $this->customers = $arr;

        return $this;
    }

}

