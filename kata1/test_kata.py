#!/usr/bin/env python3.5
import pytest
from timeit import timeit
from kata import make_kata
FIXTURES = (
    ("The Goal of this", "))()()(())()))(("),
    ("Where Each Character", "())))))))))))))))())"),
    ('Exercise is to Convert', ')()))))))))))))))(()))'),
)

@pytest.mark.parametrize("inputted, expected", FIXTURES)
def test__kata__works(inputted, expected):
    assert make_kata(inputted) == expected



def test__perfomance():
    timeit(print(make_kata("this is the Test string written in english!")))
    assert True
