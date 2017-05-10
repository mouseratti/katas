
# def to_decimal_ip(inp):
#     while inp:
#         octet = inp[:8]
#         inp = inp[9:]
#         octet = [2 * int(x) for x in octet]


# def make_kata(ipaddress):
#     ip, mask = ipaddress.split('/')
#     binary_ip = ''.join(['%8s' % bin(int(x))[2:] for x in ip.split('.')]).replace(' ', '0')
#     assert len(binary_ip) == 32
#     mask = int(mask)
#     wildcard = 32 - mask
#     binary_subnet = binary_ip[:mask] + '0' * wildcard
#     binary_broadcast = binary_ip[:mask] + '1' * wildcard
#     binary_wildcard = '0' * mask + bin(wildcard)[2:]
#     subnet = []
#     broadcast = []

#     while binary_subnet:
#         octet = binary_subnet[:8]
#         binary_subnet = binary_subnet[8:]
#         octet = sum([int(mult) * 2 **(7 - pos) for pos, mult in enumerate(octet)])
#         subnet.append(octet)
#     subnet = '.'.join([str(x) for x in subnet])

#     while binary_broadcast:
#         octet = binary_broadcast[:8]
#         binary_broadcast = binary_broadcast[8:]
#         octet = sum([int(mult) * 2 **(7 - pos) for pos, mult in enumerate(octet)])
#         broadcast.append(octet)
#     broadcast = '.'.join([str(x) for x in broadcast])
#     return subnet, broadcast


# def make_kata(ipaddr):
#     subnet = broadcast = ''
#     ip, mask = ipaddr.split("/")
#     ip = ''.join([bin(int(x) + 256)[3:] for x in ip.split(".")])
#     mask = int(mask)
#     assert len(ip) == 32, "ip is %s; len(ip) is %s" % (ip, len(ip))
#     subnet = ip[:mask] + '0' * (32 - mask)
#     broadcast = ip[:mask] + '1' * (32 - mask)
#     assert len(subnet) == len(broadcast) == 32
#     def to_decimal_ip(binary_ip):
#         iplist = [binary_ip[x:x+8] for x in range(0, len(binary_ip), 8)]
#         for position, binary_octet in enumerate(iplist):
#             octet = 0
#             for pos, symbol in enumerate(binary_octet):
#                 octet += int(symbol) * 2 ** (7 - pos)
#             iplist[position] = str(octet)
#         return '.'.join(iplist)
#     return to_decimal_ip(subnet), to_decimal_ip(broadcast)


def make_kata(ipaddr):
    ip, mask = ipaddr.split("/")
    mask = int(mask)
    ip = ''.join([bin(int(octet) + 256)[3:] for octet in ip.split('.')])
    assert len(ip) == 32
    subnet = ip[:mask] + '0' * (32 - mask)
    broadcast = ip[:mask] + '1' * (32 - mask)

    def to_decimal_ip(ip):
        iplist = [ip[x:x + 8] for x in range(0, 32, 8)]
        for position, octet in enumerate(iplist):
            iplist[position] = str(sum([int(sym) * 2 ** (7 - pos) for pos, sym in enumerate(octet)]))
        return '.'.join(iplist)
    return to_decimal_ip(subnet), to_decimal_ip(broadcast)
