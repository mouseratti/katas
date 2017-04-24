
# def to_decimal_ip(inp):
#     while inp:
#         octet = inp[:8]
#         inp = inp[9:]
#         octet = [2 * int(x) for x in octet]




def make_kata(ipaddress):
    ip, mask = ipaddress.split('/')
    binary_ip = ''.join(['%8s' % bin(int(x))[2:] for x in ip.split('.')]).replace(' ', '0')
    assert len(binary_ip) == 32
    mask = int(mask)
    wildcard = 32 - mask
    binary_subnet = binary_ip[:mask] + '0' * wildcard
    binary_broadcast = binary_ip[:mask] + '1' * wildcard
    binary_wildcard = '0' * mask + bin(wildcard)[2:]
    subnet = []
    broadcast = []

    while binary_subnet:
        octet = binary_subnet[:8]
        binary_subnet = binary_subnet[8:]
        octet = sum([int(mult) * 2 **(7 - pos) for pos, mult in enumerate(octet)])
        subnet.append(octet)
    subnet = '.'.join([str(x) for x in subnet])

    while binary_broadcast:
        octet = binary_broadcast[:8]
        binary_broadcast = binary_broadcast[8:]
        octet = sum([int(mult) * 2 **(7 - pos) for pos, mult in enumerate(octet)])
        broadcast.append(octet)
    broadcast = '.'.join([str(x) for x in broadcast])
    return subnet, broadcast
