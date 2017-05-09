import pytest
from kata import make_kata

fixtures = [
    ("192.168.0.1/24", "192.168.0.0", "192.168.0.255"),
    ("192.168.0.1/25", "192.168.0.0", "192.168.0.127"),
    ("192.168.10.129/30", "192.168.10.128", "192.168.10.131"),
    ("10.20.30.40/27", "10.20.30.32", "10.20.30.63"),
    ("10.0.0.133/8", "10.0.0.0", "10.255.255.255"),
]



@pytest.mark.parametrize("inp, subnet, broadcast", fixtures)
def test_kata3(inp, subnet, broadcast):
    assert make_kata(inp) == (subnet, broadcast)



@pytest.mark.benchmark(
    # min_rounds=30000,
    # disable_gc=True,
    # warmup=True
)
def test_benchmark(benchmark):
    benchmark(make_kata, "10.20.30.40/28")