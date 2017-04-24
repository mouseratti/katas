import pytest
from kata import make_kata

fixtures = [
    ("192.168.0.1/24", "192.168.0.0", "192.168.0.255"),
    ("192.168.0.1/25", "192.168.0.0", "192.168.0.127"),
    ("192.168.10.129/30", "192.168.10.128", "192.168.10.131"),

]



@pytest.mark.parametrize("inp, subnet, broadcast", fixtures)
def test_kata3(inp, subnet, broadcast):
    assert make_kata(inp) == (subnet, broadcast)
