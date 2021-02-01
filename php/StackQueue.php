<?php
class stack{
    protected $stack = [];

    public function push($value){
        array_push($this->stack, $value);
    }

    public function pop(){
        return array_pop($this->stack);
    }

    public function length(){
        return count($this->stack);
    }

}

class stackQueue{
    protected $enterStack;
    protected $outStack;

    public function __construct()
    {
        $this->enterStack = new stack();
        $this->outStack = new stack();
    }

    public function push($value){
        $this->enterStack->push($value);
    }

    public function pop(){
        if(!$this->outStack->length() && !$this->enterStack->length()) throw new Exception("队列为空");
        if(!$this->outStack->length()){
            $k = $this->enterStack->length();
            for ($i = 0; $i < $k; $i ++ ){
                $value = $this->enterStack->pop();
                $this->outStack->push($value);
            }
        }

        return $this->outStack->pop();
    }
}

function run(){
    $stackQueue = new stackQueue();
    $stackQueue->push(1);
    $stackQueue->push(3);
    $stackQueue->push(5);
    $stackQueue->push(7);
    $stackQueue->push(8);
    for ($i = 0; $i < 6; $i ++ ){
        var_dump("out: " . $stackQueue->pop());
    }
}

run();