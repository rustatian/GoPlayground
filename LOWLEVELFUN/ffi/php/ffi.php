<?php
// create FFI object, loading libc and exporting function ()
// https://www.php.net/manual/en/ffi.examples-basic.php

//void print_string(char *s);

$ffi = FFI::cdef(
    "int fib(int n);",
    "/home/valery/Projects/repo/GoPlayground/LOWLEVELFUN/ffi/php/libimage_ffi_php.so");

$time_start = microtime(true);

for ($i=0; $i < 1000000; $i++) {
	$ffi->fib(12);
}

echo '[Rust] Release execution time:' . (microtime(true) - $time_start).PHP_EOL;



// $time_start = microtime(true);


// echo $v;
// echo '[Rust] Release execution time:' . (microtime(true) - $time_start).PHP_EOL;