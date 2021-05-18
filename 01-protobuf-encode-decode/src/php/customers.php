<?php

require "vendor/autoload.php";

$customer1 = (new Customer())
    ->setId(1)
    ->setName('John');

$customer2 = (new Customer())
    ->setId(2)
    ->setName('Bob');

$customers = new Customers();
$customers->setCustomers([$customer1, $customer2]);

$bytes = $customers->serializeToString();

$customersFromBytes = new Customers();
$customersFromBytes->mergeFromString($bytes);

echo $customersFromBytes->serializeToJsonString();
