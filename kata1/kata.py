
# def make_kata(input_string):
#     """Test finished in 2.167048931121826 seconds """
#     input_string = input_string.lower()
#     output = ''
#     inputted_list = list(input_string)
#     inputted_set = set(inputted_list)
#     for symbol in inputted_set:
#         inputted_list.remove(symbol)
#     for symbol in input_string:
#         output += ')' if symbol in inputted_list else '('
#     return output

# def make_kata(inputted):
#     """ Test finished in 0.9294266700744629 seconds """
#     d = {}
#     inputted = inputted.lower()
#     outputted = ''
#     for symbol in inputted:
#         d[symbol] = ')' if symbol in d else '('
#     for symbol in inputted:
#         outputted += d[symbol]
#     return outputted

# def make_kata(inputted):
#     d = {}
#     def f(stri):
#         length = len(stri)
#         if length == 2:
#             stri = stri.lover()
#             if
#             if stri[0] == stri[1]:
#                 d[stri[0]] = ')'

#     d = set()
#     outputted = ''
#     outputted_list = []
#     for symbol in inputted:
#         symbol = symbol.lower()
#         outputted_list.append() = ')' if symbol in d else '('
#     for symbol in inputted:
#         outputted += d[symbol]
#     return outputted


# def make_kata(inputted):
#     """Test finished in 2.322965145111084 seconds"""
#     d = {}
#     inputted = inputted.lower()
#     inputted_list = list(inputted)
#     while inputted_list:
#         symbol = inputted_list.pop(0)
#         if symbol not in d:
#             d[symbol] = ')' if symbol in inputted_list else '('
#     for symbol in d:
#         inputted = inputted.replace(symbol, d[symbol])
#     return inputted

# def make_kata(inputted):
#     """Test finished in 2.30079984664917 seconds"""
#     d = {}
#     inputted = inputted.lower()
#     inputted_list = list(inputted)
#     while inputted_list:
#         symbol = inputted_list.pop(0)
#         if symbol in d:
#             continue
#         d[symbol] = ')' if symbol in inputted_list else '('
#     for symbol in d:
#         inputted = inputted.replace(symbol, d[symbol])
#     return inputted

# def make_kata(inputted):
#     """Test finished in 1.7508909702301025 seconds"""
#     d = set()
#     inputted = inputted.lower()
#     while inputted[0] not in ('(', ')'):
#         symbol = inputted[0]
#         inputted = inputted[1:]
#         if symbol in d:
#             inputted += ')'
#             continue
#         if symbol in inputted:
#             d.add(symbol)
#             inputted += ')'
#             continue
#         inputted += '('
#     return inputted

# def make_kata(inputted):
#     """Test finished in 1.748342514038086 seconds"""
#     d = ''
#     inputted = inputted.lower()
#     while inputted[0] not in ('(', ')'):
#         symbol = inputted[0]
#         inputted = inputted[1:]
#         if symbol in d:
#             inputted += ')'
#             continue
#         if symbol in inputted:
#             d += symbol
#             inputted += ')'
#             continue
#         inputted += '('
#     return inputted

# def make_kata(inputted):
#     """Test finished in 1.7696144580841064 seconds"""
#     d = ''
#     inputted = inputted.lower()
#     symbol = inputted[0]
#     while  symbol not in ('(', ')'):
#         inputted = inputted[1:]
#         if symbol in inputted:
#             d += symbol
#             inputted += ')'
#             symbol = inputted[0]
#             continue
#         if symbol in d:
#             inputted += ')'
#             symbol = inputted[0]
#             continue
#         inputted += '('
#         symbol = inputted[0]
#     return inputted


def make_kata(inputted):
    """Test finished in 2.796117067337036 seconds"""
    inputted = inputted.lower()
    inputted_list = list(inputted)
    def gen(li):
        nonuniq = []
        while li:
            symbol = li.pop(0)
            if symbol in nonuniq:
                out = ')'
            elif symbol in li:
                out = ')'
                nonuniq.append(symbol)                
            else:
                out = '('
            yield out
    return ''.join(_ for _ in gen(inputted_list))
