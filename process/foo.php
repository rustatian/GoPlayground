<?php

echo getmypid();
echo "-initial\n";
$pid = pcntl_fork();
if ($pid == -1) {
     die('could not fork');
} else if ($pid) {
    echo $pid;
    echo "-child\n";

    echo getmypid();
    echo "-parent\n";
     // we are the parent
     pcntl_wait($status); //Protect against Zombie children
} else {
    echo getmypid();
    echo "-child\n";
     // we are the child
}

?>
