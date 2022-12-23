<?php

echo getmypid();
echo "-initial\n";

for ($i = 0; $i <= 10000; ++$i) {
     sleep(1);
     echo $i;
}
// $pid = pcntl_fork();
// if ($pid == -1) {
//      die('could not fork');
// } else if ($pid) {
//     echo $pid;
//     echo "-child\n";

//     echo getmypid();
//     echo "-parent\n";
//      // we are the parent
//      pcntl_wait($status); //Protect against Zombie children
// } else {
//     echo getmypid();
//     echo "-child\n";
//      // we are the child
// }

?>
