#!/usr/bin/env python3.5
import pytest
from timeit import timeit
from kata import make_kata
import time
FIXTURES = (
    ("The Goal of this", "))()()(())()))(("),
    ("Where Each Character", "())))))))))))))))())"),
    ('Exercise is to Convert', ')()))))))))))))))(()))'),
)

@pytest.mark.parametrize("inputted, expected", FIXTURES)
def test__kata__works(inputted, expected):
    assert make_kata(inputted) == expected



def test__performance():
    start = time.time()
    i = 0
    while i <= 100000:
        make_kata("This is the best string I ever Watched")
        i += 1
    finish = time.time()
    print("Test finished in %s seconds " % (finish - start) )
    assert False
